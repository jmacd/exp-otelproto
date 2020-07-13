package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary_async"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/http11"
)

func main() {
	options := core.Options{}

	// ballastSizeBytes := uint64(4000) * 1024 * 1024
	ballastSizeBytes := 0
	ballast := make([]byte, ballastSizeBytes)

	protocol := flag.String("protocol", "",
		"protocol to benchmark (unary,unaryasync,http11,http11conc)",
	)

	flag.IntVar(&options.Batches, "batches", 100, "total batches to send")
	flag.IntVar(&options.MetricsPerBatch, "metricsperbatch", 100, "metrics per batch")
	flag.IntVar(&options.PointsPerMetric, "pointspermetric", 1, "points per metric")
	flag.IntVar(&options.LabelsPerMetric, "labelspermetric", 10, "labels per metric")
	flag.IntVar(&options.LabelNumValues, "labelnumvalues", 3, "number of distinct values")

	var aggName = flag.String("aggregation", "", "aggregation name")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memprofile = flag.String("memprofile", "", "write mem profile to file")

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	// if *memprofile != "" {
	// 	f, err := os.Create(*memprofile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pprof.StartMEMProfile(f)
	// 	defer pprof.StopMEMProfile()
	// }

	switch *aggName {
	case "mmlsc":
		options.Aggregation = core.MMLSC
	case "float64":
		options.Aggregation = core.FLOAT64
	case "int64":
		options.Aggregation = core.INT64
	default:
		log.Fatal("unrecognized aggregation: ", *aggName)
	}

	switch *protocol {
	// case "unary":
	// 	benchmarkGRPCUnary(options)
	// case "unaryasync":
	// 	benchmarkGRPCUnaryAsync(options)
	case "http11":
		benchmarkHttp11(options, 1)
	case "http11conc":
		benchmarkHttp11(options, 10)
	default:
		flag.Usage()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	runtime.KeepAlive(ballast)
}

func benchmarkGRPCUnary(options core.Options) {
	benchmarkImpl(
		"OTLP/GRPC-Unary/Sequential",
		options,
		func() core.Client { return &grpc_unary.Client{} },
		func() core.Server { return &grpc_unary.Server{} },
		func() core.MetricGenerator { return otlp.NewGenerator() },
	)
}

func benchmarkGRPCUnaryAsync(options core.Options) {
	benchmarkImpl(
		"OTLP/GRPC-Unary/Concurrent",
		options,
		func() core.Client { return &grpc_unary_async.Client{} },
		func() core.Server { return &grpc_unary_async.Server{} },
		func() core.MetricGenerator { return otlp.NewGenerator() },
	)
}

func benchmarkHttp11(options core.Options, concurrency int) {
	benchmarkImpl(
		"OTLP/HTTP1.1/"+strconv.Itoa(concurrency),
		options,
		func() core.Client { return &http11.Client{Concurrency: concurrency} },
		func() core.Server { return &http11.Server{} },
		func() core.MetricGenerator { return otlp.NewGenerator() },
	)
}

func benchmarkImpl(
	name string,
	options core.Options,
	clientFactory func() core.Client,
	serverFactory func() core.Server,
	generatorFactory func() core.MetricGenerator,
) {
	cpuSecs, wallSecs := core.BenchmarkLocalDelivery(
		clientFactory,
		serverFactory,
		generatorFactory,
		options,
	)

	fmt.Printf("%-28s %7d metrics|%5.1f cpusec|%5.1f wallsec|%7.1f batches/cpusec|%8.1f batches/wallsec|%5.1f cpuμs/metric\n",
		name,
		options.Batches*options.MetricsPerBatch,
		cpuSecs,
		wallSecs,
		float64(options.Batches)/cpuSecs,
		float64(options.Batches)/wallSecs,
		cpuSecs*1e6/float64(options.Batches*options.MetricsPerBatch),
	)
}
