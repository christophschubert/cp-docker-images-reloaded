#!/usr/bin/env bash

#fail when a subprogram fails
set -e

# Usage:
function exit_if_not_set {
  param=$1
  if [[ -z ${!param} ]]
  then
    echo "  Required environment variable $param not set"
    exit 1
  fi
}

# Usage
# kafka_ready numBrokers timeout bootstrapServer pathToConfig
function kafka_ready {
  #TODO: make values configurable via env vars
  if java $KAFKA_OPTS -cp "$CUB_CLASSPATH" "io.confluent.admin.utils.cli.KafkaReadyCommand" --bootstrap-servers $3 --config $4 $1 $2
  then
    echo "Kafka ready: found at least $1 broker(s) at $3"
  else
    exit 1
  fi
}