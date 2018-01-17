package main

import "github.com/prometheus/client_golang/prometheus"

// AkamaiCollector
type AkamaiCollector struct {}

var (
	// declare Prometheus metric descriptions
)

// Describe implements a method from the Collector interface created by Prometheus
func (c *AkamaiCollector) Describe(ch chan<- *prometheus.Desc) {

}

// Collect implements a method from the Collector interface created by Prometheus
func (c *AkamaiCollector) Collect(ch chan<- prometheus.Metric) {

}
