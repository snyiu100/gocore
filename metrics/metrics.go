package metrics

import "time"

// MetricsRecorder global instance.
var Metrics MetricsRecorder

// Public interface for recording metrics.
type MetricsRecorder interface {
	IncrementCount(metricName string)
	IncrementCountBy(metricName string, amount int)
	MeasureSince(metricName string, since time.Time)
	SetGauge(metricName string, val float32)
	SetPrefix(prefix string)
}

// Package-level default initialization of the Metrics global.
// Initializes it to a no-op implementation;
// later calls can replace it by calling SetupMetrics with a real stats location.
func init() {
	Metrics = &NoopRecorder{}
}

// Public initialization function to initialize the Metrics global.
// If you're using metrics, this should be called before any goroutines
// using them are started.
//
// (If you don't care about metrics, you don't need to call this function;
// nothing will break, since a no-op metrics sink is used by default.)
func SetMetricsGlobal(recorder MetricsRecorder) {
	Metrics = recorder
}

// Increment Count by 1 for Metric by name
func IncrementCount(metricName string) {
	Metrics.IncrementCount(metricName)
}

// Increment Count by amount for Metric by name
func IncrementCountBy(metricName string, amount int) {
	Metrics.IncrementCountBy(metricName, amount)
}

// Measure Time since given for Metric by name
func MeasureSince(metricName string, since time.Time) {
	Metrics.MeasureSince(metricName, since)
}

// Gauge value for Metric by name
func SetGauge(metricName string, val float32) {
	Metrics.SetGauge(metricName, val)
}

// Set Prefix for all Metrics collected
func SetPrefix(prefix string) {
	Metrics.SetPrefix(prefix)
}