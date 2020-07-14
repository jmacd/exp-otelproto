#!/usr/bin/env bash

# Set MULTIPLIER to 1 for quick results and to 100 for more stable results.
MULTIPLIER=50

echo ====================================================================================
echo Legend:
echo "http/otlp/N     - OTLP baseline, ProtoBuf,HTTP 1.1, N concurrent requests."
echo "http/experiment/N  - OTLP experiment, ProtoBuf,HTTP 1.1, N concurrent requests."
echo

benchmark() {
    ./benchmark -protocol "${1}" -encoding "${2}" -batches=${BATCHES} -metricsperbatch=${METRICSPERBATCH} -labelspermetric=${LABELSPERMETRIC} -aggregation "${3}"
}

benchmark_all() {
    echo ${BATCHES} $1 batches, ${METRICSPERBATCH} metrics per batch, ${LABELSPERMETRIC} labels per metric
    benchmark http11 otlp int64
    benchmark http11conc otlp int64
    benchmark http11 otlp float64
    benchmark http11conc otlp float64
    benchmark http11 otlp mmlsc
    benchmark http11conc otlp mmlsc

    benchmark http11 experiment int64
    benchmark http11conc experiment int64
    benchmark http11 experiment float64
    benchmark http11conc experiment float64
    benchmark http11 experiment mmlsc
    benchmark http11conc experiment mmlsc
    echo
}

./beforebenchmarks.sh

echo

cd bin

let BATCHES=6400*MULTIPLIER
METRICSPERBATCH=1
LABELSPERMETRIC=10
benchmark_all nano

let BATCHES=1600*MULTIPLIER
METRICSPERBATCH=10
LABELSPERMETRIC=10
benchmark_all tiny


let BATCHES=800*MULTIPLIER
METRICSPERBATCH=100
LABELSPERMETRIC=10
benchmark_all small


let BATCHES=80*MULTIPLIER
METRICSPERBATCH=500
LABELSPERMETRIC=10
benchmark_all large

let BATCHES=10*MULTIPLIER
METRICSPERBATCH=5000
LABELSPERMETRIC=10
benchmark_all "very large"

echo ====================================================================================

cd ..

./afterbenchmarks.sh
