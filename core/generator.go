package core

// SpanGenerator allows to generate a ExportRequest containing a batch of spans.
type SpanGenerator interface {
	GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) ExportRequest
}

// MetricGenerator allows to generate a ExportRequest containing a batch of metrics.
type MetricGenerator interface {
	GenerateMetricBatch(
		metricsPerBatch int,
		pointsPerMetric int,
		labelsPerMetric int,
		labelNumValues float64,
		aggregation Aggregation,
	) ExportRequest
}

// LogGenerator allows to generate a ExportRequest containing a batch of log.
type LogGenerator interface {
	GenerateLogBatch(logsPerBatch int, attrsPerLog int) ExportRequest
}

type Generator interface {
	SpanGenerator
	MetricGenerator
	LogGenerator
}

type LabelMaker struct {
}
type Label struct {
	Key   string
	Value string
}

func (l *LabelMaker) Get() []Label {
	return nil // @@@
}

func NewLabelMaker(labelsPerMetric int, labelNumValues float64) *LabelMaker {
	return &LabelMaker{}
}
