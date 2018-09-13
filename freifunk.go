package main

type nodes struct {
	Meta struct {
		Timestamp string `json:"timestamp"`
	} `json:"meta"`
	Version int `json:"version"`
	Nodes   []struct {
		Lastseen string `json:"lastseen"`
		Nodeinfo struct {
			Owner struct {
				Contact string `json:"contact"`
			} `json:"owner"`
			Location struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"location"`
			NodeID   string `json:"node_id"`
			Software struct {
				Firmware struct {
					Release string `json:"release"`
				} `json:"firmware"`
			} `json:"software"`
			System struct {
				Role     string `json:"role"`
				SiteCode string `json:"site_code"`
			} `json:"system"`
			Network struct {
				Addresses []string `json:"addresses"`
				Mac       string   `json:"mac"`
			} `json:"network"`
			Hostname string `json:"hostname"`
			Hardware struct {
				Model string `json:"model"`
			} `json:"hardware"`
		} `json:"nodeinfo"`
		Flags struct {
			Online  bool `json:"online"`
			Gateway bool `json:"gateway"`
		} `json:"flags"`
		Firstseen  string `json:"firstseen"`
		Statistics struct {
			MemoryUsage float64 `json:"memory_usage"`
			Clients     int     `json:"clients"`
			RootfsUsage float64 `json:"rootfs_usage"`
			Uptime      float64 `json:"uptime"`
			Loadavg     float64 `json:"loadavg"`
		} `json:"statistics"`
	} `json:"nodes"`
}
