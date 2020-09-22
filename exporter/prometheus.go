package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// Describe - loops through the API metrics and passes them to prometheus.Describe
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	for _, m := range e.APIMetrics {
		ch <- m
	}

}

// Collect function, called on by Prometheus Client library
// This function is called when a scrape is peformed on the /metrics page
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	client := e.newClient()

	// Orgs
	// TODO - What are we doing
	e.gatherOrgMetrics(client)

	// Users

	// Repos

	// data, rate, err := e.gatherData(client)
	// if err != nil {
	// 	log.Errorf("Error gathering Data from github API: %v", err)
	// 	return
	// }

	// Set prometheus gauge metrics using the data gathered
	err := e.processMetrics(ch)

	if err != nil {
		log.Error("Error Processing Metrics", err)
		return
	}

	log.Info("All Metrics successfully collected")

}
