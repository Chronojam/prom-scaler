package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chronojam/prom-scaler/config"
	"github.com/chronojam/prom-scaler/scaler"

	_ "github.com/chronojam/prom-scaler/scaler/kubernetes"
)

var tt map[string]scaler.ScalableResource

func main() {
	config, err := config.Load("/etc/scaler/config/config.yaml")
	if err != nil {
		panic(err)
	}

	tt = map[string]scaler.ScalableResource{}

	// Go through each scalar config and initialize them with the appropriate driver
	// store the result in a map we can use to look up later to scale with.
	for _, s := range config.Scalars {
		fmt.Println("Getting s config.")
		fmt.Println(scaler.Drivers)
		fmt.Println(s.Type)
		if init, ok := scaler.Drivers[s.Type]; ok {
			fmt.Println("Matched type to driver.")
			r, err := init(s)
			if err != nil {
				panic(err)
			}
			fmt.Println(r)
			fmt.Println(r.Names())

			// Each scalar can have many names
			for _, n := range r.Names() {
				fmt.Println("Populating tt..")
				tt[n] = r
				fmt.Println(tt)
			}
		}
	}

	// WebServer to recieve post's
	fmt.Println("Serving on 8080")
	http.HandleFunc("/alerts", scaleHandler)
	http.ListenAndServe(":8080", nil)
}

func scaleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Scale.")
	fmt.Println("tt: ", tt)
	decoder := json.NewDecoder(r.Body)
	var pa PrometheusAlert
	err := decoder.Decode(&pa)
	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()
	for _, a := range pa.Alerts {
		if n, ok := a.Labels["scale_name"]; ok {
			name, _ := n.(string)
			fmt.Println(name)
			if stype, ok := a.Labels["scale_type"]; ok {
				if stype == "up" {
					fmt.Println(tt[name].Scale(1))
				}
				if stype == "down" {
					fmt.Println(tt[name].Scale(-1))
				}
			}
		}
	}
	fmt.Println(pa.Alerts)
}
