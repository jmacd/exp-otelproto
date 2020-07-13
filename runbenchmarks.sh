#!/usr/bin/env bash

# Set MULTIPLIER to 1 for quick results and to 100 for more stable results.
MULTIPLIER=10

echo ====================================================================================
echo Legend:
echo "OTLP/GRPC-Unary/Sequential  - OTLP, Unary, sequential. One request per batch, load balancer friendly, with ack"
echo "OTLP/GRPC-Unary/Concurrent  - OTLP, Unary, 10 concurrent requests, load balancer friendly, with ack"
echo "OTLP/HTTP1.1/N              - OTLP ProtoBuf,HTTP 1.1, N concurrent requests. Load balancer friendly."
echo

benchmark() {
    ./benchmark -protocol ${1} -batches=${BATCHES} -metricsperbatch=${METRICSPERBATCH} -labelspermetric=${LABELSPERMETRIC} -aggregation int64 -
}

benchmark_all() {
    echo ${BATCHES} $1 batches, ${METRICSPERBATCH} metrics per batch, ${LABELSPERMETRIC} attrs per metric
    #benchmark sapm
    benchmark http11
    benchmark http11conc
    #benchmark wsstreamsync
    #benchmark wsstreamasync
    #benchmark wsstreamasyncconc
    #benchmark wsstreamasynczlib
    #benchmark unary
    #benchmark unaryasync
    #benchmark streamlbasync
    #benchmark streamlbconc
    #benchmark opencensus
    #benchmark ocack
    #benchmark streamsync
    #benchmark streamlbalwayssync
    #benchmark streamlbtimedsync
    #benchmark streamlbsrv
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
