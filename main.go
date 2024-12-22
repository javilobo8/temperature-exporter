package main

import (
	"fmt"
	"log"
	"net/http"
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
	temperatures := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "sensor_temperatures",
		Help: "Sensor temperatures",
	}, []string{"sensor"})

	registry := prometheus.NewRegistry()

	registry.MustRegister(temperatures)

	go func() {
		for {
			s, _ := sensors.SensorsTemperatures()
			for _, v := range s {
				temperatures.WithLabelValues(v.SensorKey).Set(v.Temperature)
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
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
