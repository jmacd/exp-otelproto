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
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlpexperiment"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/http11"
)

func main() {
	options := core.Options{}

	// ballastSizeBytes := uint64(4000) * 1024 * 1024
	ballastSizeBytes := 0
	ballast := make([]byte, ballastSizeBytes)

	protocol := flag.String("protocol", "", "protocol to benchmark (http11,http11conc)")
	encoding := flag.String("encoding", "", "encoding to benchmark (otlp,experiment)")

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

	var gen func() core.Generator
	switch *encoding {
	case "otlp":
		gen = otlp.NewGenerator
	case "experiment":
		gen = otlpexperiment.NewGenerator
	default:
		log.Fatal("unrecognized encoding: ", *encoding)
	}

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
	name := fmt.Sprint(*protocol, "/", *encoding, "/", *aggName)

	switch *protocol {
	case "http11":
		benchmarkHttp11(name, options, 1, gen)
	case "http11conc":
		benchmarkHttp11(name, options, 10, gen)
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

func benchmarkHttp11(name string, options core.Options, concurrency int, gen func() core.Generator) {
	benchmarkImpl(
		name+"/"+strconv.Itoa(concurrency),
		options,
		func() core.Client { return &http11.Client{Concurrency: concurrency} },
		func() core.Server { return &http11.Server{} },
		gen,
	)
}

func benchmarkImpl(
	name string,
	options core.Options,
	clientFactory func() core.Client,
	serverFactory func() core.Server,
	generatorFactory func() core.Generator,
) {
	cpuSecs, wallSecs := core.BenchmarkLocalDelivery(
		clientFactory,
		serverFactory,
		generatorFactory,
		options,
	)

	fmt.Printf("%-28s %7d points|%5.1f cpusec|%5.1f wallsec|%7.1f batches/cpusec|%8.1f batches/wallsec|%5.1f cpuμs/point\n",
		name,
		options.Batches*options.MetricsPerBatch*options.PointsPerMetric,
		cpuSecs,
		wallSecs,
		float64(options.Batches)/cpuSecs,
		float64(options.Batches)/wallSecs,
		cpuSecs*1e6/float64(options.Batches*options.MetricsPerBatch*options.PointsPerMetric),
	)
}
