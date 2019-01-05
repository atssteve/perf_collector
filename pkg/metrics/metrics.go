package metrics

// Metric provides an interface for all of the metrics we have collected.
type Metric interface {
	String() string
}
