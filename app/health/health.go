package health

import (
	"fmt"
	"log"
	"runtime"
)


const (
	StatusHealthy = "healthy"
	StatusError   = "error"

	MemoryThresh  = 2*1024*1024*1024 // bytes
)

var indicators = []healthIndicator{PingHealthIndicator, MemoryIndicator}

type HealthStatus struct {
	Name	   string
	Status     string
	Message    string
}

type healthIndicator func() HealthStatus


func PingHealthIndicator() HealthStatus {
	log.Default().Print("Health check: Healthy")
	return HealthStatus{"Ping", StatusHealthy, ""}
}

func MemoryIndicator() HealthStatus {
	status := HealthStatus{}
	status.Name = "Memory"
	status.Status = StatusHealthy

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	if m.Sys > MemoryThresh {
		status.Status  = StatusError
		status.Message = "Memory usage crossed threshold"
		
		log.Default().Print(status.Message + " [Memory: " + fmt.Sprintf("%d", m.Sys) + " bytes ]")
	}

	return status
}

func Health() ([]HealthStatus, bool) {
	var results []HealthStatus = make([]HealthStatus, 0, len(indicators))
	var err bool = false
	
	for _, status := range indicators {
		result := status()
		if result.Status == StatusError {
			err = true
		}
		results = append(results, result)
	}

	return results, err
}

