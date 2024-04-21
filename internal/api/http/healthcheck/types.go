// Package healthcheck provides primitives to interact with the openapi HTTP API.
package healthcheck

// Defines values for SuccessProbe.
const (
	Healthy    string = "Healthy"
	Ready      string = "Ready"
	NotHealthy string = "Not healthy"
	NotReady   string = "Not ready"
)
