package main

type PrometheusAlert struct {
	Version string
	Status  string
	Alerts  []struct {
		Labels      map[string]interface{}
		Annotations map[string]interface{}
		StartsAt    string
		EndsAt      string
	}
}
