package collectors

func init() {
	registerCollector("meminfo", NewMemInfoCollector)
}

type memInfoCollector struct {
	Collector string
}

func NewMemInfoCollector() Collector {
	return &memInfoCollector{
		Collector: "meminfo",
	}
}

func (m *memInfoCollector) Update() {
	GetMemInfo()
}
