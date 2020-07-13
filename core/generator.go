package core

import (
	"fmt"
	"math/rand"
)

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
		labelNumValues int,
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
	rnd    *rand.Rand
	keys   []string
	values [][]string
}

type Label struct {
	Key   string
	Value string
}

func (lm *LabelMaker) Get() []Label {
	l := make([]Label, len(lm.keys))
	for i := 0; i < len(l); i++ {
		j := lm.rnd.Intn(len(lm.values[i]))
		l[i] = Label{
			Key:   lm.keys[i],
			Value: lm.values[i][j],
		}
	}
	return l
}

func NewLabelMaker(labelsPerMetric int, labelNumValues int) *LabelMaker {
	keys := make([]string, labelsPerMetric)
	values := make([][]string, labelsPerMetric)
	for i := range keys {
		keys[i] = fmt.Sprint("key-", i)
		values[i] = make([]string, labelNumValues)
		for j := range values[i] {
			values[i][j] = fmt.Sprint("value-", i, "-", j)
		}
	}
	return &LabelMaker{
		rnd:    rand.New(rand.NewSource(13339)),
		keys:   keys,
		values: values,
	}
}
