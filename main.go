package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v4/sensors"
)

func HandleCollect(gauge prometheus.Gauge) {
	s, _ := sensors.SensorsTemperatures()

	for _, v := range s {
		fmt.Println(v.SensorKey, v.Temperature)
	}
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	temperatures := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "sensor_temperatures",
		Help: "Sensor temperatures",
	}, []string{"hostname", "sensor"})

	registry := prometheus.NewRegistry()

	registry.MustRegister(temperatures)

	go func() {
		for {
			s, _ := sensors.SensorsTemperatures()
			for _, v := range s {
				temperatures.WithLabelValues(hostname, v.SensorKey).Set(v.Temperature)
			}
			time.Sleep(30 * time.Second)
		}
	}()

	// Expose /metrics HTTP endpoint using the created custom registry.
	http.Handle(
		"/metrics", promhttp.HandlerFor(
			registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			}),
	)

	// To test: curl -H 'Accept: application/openmetrics-text' localhost:8080/metrics
	log.Fatalln(http.ListenAndServe(":9101", nil))
}
