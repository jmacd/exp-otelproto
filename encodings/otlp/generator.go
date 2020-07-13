package otlp

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	otlpmetriccol "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/metrics/v1"
	otlpcommon "github.com/open-telemetry/opentelemetry-proto/gen/go/common/v1"
	otlpmetric "github.com/open-telemetry/opentelemetry-proto/gen/go/metrics/v1"
	otlpresource "github.com/open-telemetry/opentelemetry-proto/gen/go/resource/v1"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

// Generator allows to generate a ExportRequest.
type Generator struct {
	random     *rand.Rand
	tracesSent uint64
	spansSent  uint64
}

func NewGenerator() *Generator {
	return &Generator{
		random: rand.New(rand.NewSource(99)),
	}
}

func (g *Generator) genRandByteString(len int) string {
	b := make([]byte, len)
	for i := range b {
		b[i] = byte(g.random.Intn(10) + 33)
	}
	return string(b)
}

func GenResource() *otlpresource.Resource {
	return &otlpresource.Resource{
		Attributes: []*otlpcommon.KeyValue{
			&otlpcommon.KeyValue{
				Key: "StartTimeUnixnano",
				Value: &otlpcommon.AnyValue{
					Value: &otlpcommon.AnyValue_IntValue{IntValue: 12345678},
				},
			},
			&otlpcommon.KeyValue{
				Key: "Pid",
				Value: &otlpcommon.AnyValue{
					Value: &otlpcommon.AnyValue_IntValue{IntValue: 1234},
				},
			},
			&otlpcommon.KeyValue{
				Key: "HostName",
				Value: &otlpcommon.AnyValue{
					Value: &otlpcommon.AnyValue_StringValue{StringValue: "fakehost"},
				},
			},
			&otlpcommon.KeyValue{
				Key: "ServiceName",
				Value: &otlpcommon.AnyValue{
					Value: &otlpcommon.AnyValue_StringValue{StringValue: "generator"},
				},
			},
		},
	}
}

func (g *Generator) GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	return nil
}

func (g *Generator) GenerateLogBatch(logsPerBatch int, attrsPerLog int) core.ExportRequest {
	return nil
}

func GenInt64Timeseries(startTime time.Time, offset int, lmaker *core.LabelMaker, pointsPerMetric int) []*otlpmetric.Int64DataPoint {
	var timeseries []*otlpmetric.Int64DataPoint
	for j := 0; j < 5; j++ {
		var points []*otlpmetric.Int64DataPoint

		for k := 0; k < pointsPerMetric; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))

			point := otlpmetric.Int64DataPoint{
				TimeUnixNano: pointTs,
				Value:        int64(offset * j * k),
				Labels:       otlpLabels(lmaker.Get()),
			}

			if k == 0 {
				point.StartTimeUnixNano = pointTs
			}

			points = append(points, &point)
		}

		timeseries = append(timeseries, points...)
	}

	return timeseries
}

func otlpLabels(labels []core.Label) []*otlpcommon.StringKeyValue {
	r := make([]*otlpcommon.StringKeyValue, len(labels))
	for i, l := range labels {
		r[i] = &otlpcommon.StringKeyValue{
			Key:   l.Key,
			Value: l.Value,
		}
	}
	return r
}

func GenFloat64Timeseries(startTime time.Time, offset int, lmaker *core.LabelMaker, pointsPerMetric int) []*otlpmetric.DoubleDataPoint {
	var timeseries []*otlpmetric.DoubleDataPoint
	for j := 0; j < 5; j++ {
		var points []*otlpmetric.DoubleDataPoint

		for k := 0; k < pointsPerMetric; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))

			point := otlpmetric.DoubleDataPoint{
				TimeUnixNano: pointTs,
				Value:        float64(offset * j * k),
				Labels:       otlpLabels(lmaker.Get()),
			}

			if k == 0 {
				point.StartTimeUnixNano = pointTs
			}

			points = append(points, &point)
		}

		timeseries = append(timeseries, points...)
	}

	return timeseries
}

func genInt64Gauge(startTime time.Time, i int, lmaker *core.LabelMaker, pointsPerMetric int) *otlpmetric.Metric {
	descr := GenMetricDescriptor(i)

	metric1 := &otlpmetric.Metric{
		MetricDescriptor: descr,
		Int64DataPoints:  GenInt64Timeseries(startTime, i, lmaker, pointsPerMetric),
	}

	return metric1
}

func genFloat64Gauge(startTime time.Time, i int, lmaker *core.LabelMaker, pointsPerMetric int) *otlpmetric.Metric {
	descr := GenMetricDescriptor(i)

	metric1 := &otlpmetric.Metric{
		MetricDescriptor: descr,
		DoubleDataPoints: GenFloat64Timeseries(startTime, i, lmaker, pointsPerMetric),
	}

	return metric1
}

func GenMetricDescriptor(i int) *otlpmetric.MetricDescriptor {
	descr := &otlpmetric.MetricDescriptor{
		Name:        "metric" + strconv.Itoa(i),
		Description: "some description: " + strconv.Itoa(i),
		Type:        otlpmetric.MetricDescriptor_INT64,
	}
	return descr
}

func genMMLSC(startTime time.Time, i int, lmaker *core.LabelMaker, pointsPerMetric int) *otlpmetric.Metric {
	descr := GenMetricDescriptor(i)
	descr.Type = otlpmetric.MetricDescriptor_SUMMARY

	var timeseries2 []*otlpmetric.SummaryDataPoint
	for j := 0; j < 1; j++ {
		var points []*otlpmetric.SummaryDataPoint

		for k := 0; k < pointsPerMetric; k++ {
			pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))
			val := float64(i * j * k)
			point := otlpmetric.SummaryDataPoint{
				TimeUnixNano: pointTs,
				Count:        3,
				Sum:          3 * val,
				Last:         val,
				Min:          val - 1,
				Max:          val + 1,
			}
			if k == 0 {
				point.StartTimeUnixNano = pointTs
			}
			points = append(points, &point)
		}

		timeseries2 = append(timeseries2, points...)
	}

	metric2 := &otlpmetric.Metric{
		MetricDescriptor:  descr,
		SummaryDataPoints: timeseries2,
	}

	return metric2
}

func (g *Generator) GenerateMetricBatch(
	metricsPerBatch int,
	pointsPerMetric int,
	labelsPerMetric int,
	labelNumValues int,
	aggregation core.Aggregation,
) core.ExportRequest {

	lmaker := core.NewLabelMaker(labelsPerMetric, labelNumValues)

	il := &otlpmetric.InstrumentationLibraryMetrics{}
	batch := &otlpmetriccol.ExportMetricsServiceRequest{
		ResourceMetrics: []*otlpmetric.ResourceMetrics{
			{
				Resource:                      GenResource(),
				InstrumentationLibraryMetrics: []*otlpmetric.InstrumentationLibraryMetrics{il},
			},
		},
	}

	for i := 0; i < metricsPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		switch aggregation {
		case core.INT64:
			il.Metrics = append(il.Metrics, genInt64Gauge(startTime, i, lmaker, pointsPerMetric))
		case core.FLOAT64:
			il.Metrics = append(il.Metrics, genFloat64Gauge(startTime, i, lmaker, pointsPerMetric))
		case core.MMLSC:
			il.Metrics = append(il.Metrics, genMMLSC(startTime, i, lmaker, pointsPerMetric))
		// case core.HISTOGRAM:
		//      il.Metrics = append(il.Metrics,
		//          genHistogram(startTime, i, lmaker, pointsPerMetric))
		default:
			log.Fatal("impossible")
		}
	}

	return batch
}

func timeToTimestamp(t time.Time) *timestamp.Timestamp {
	nanoTime := t.UnixNano()
	return &timestamp.Timestamp{
		Seconds: nanoTime / 1e9,
		Nanos:   int32(nanoTime % 1e9),
	}
}
