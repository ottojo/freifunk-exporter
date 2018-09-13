package main

import (
	"encoding/json"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("listen-address", ":9531", "The address to listen on for HTTP requests.")
var nodesURL = flag.String("nodes-url", "https://map.freifunk-ulm.de/data/nodes.json", "The location of nodes.json")

var clientsConnected = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "freifunk_clients",
		Help: "Current number of clients connected",
	},
	[]string{"node_id"},
)

func init() {
	flag.Parse()
	prometheus.MustRegister(clientsConnected)
}

func main() {

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

				resp, err := http.Get(*nodesURL)
				if err != nil {
					log.Println(err)
					continue
				}

				var n nodes
				err = json.NewDecoder(resp.Body).Decode(&n)
				if err != nil {
					log.Println(err)
					continue
				}

				for _, node := range n.Nodes {
					clientsConnected.WithLabelValues(node.Nodeinfo.NodeID).Set(float64(node.Statistics.Clients))
				}


			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
