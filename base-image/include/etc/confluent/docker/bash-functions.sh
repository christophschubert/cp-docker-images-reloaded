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

function warn_jmx_rmi_port {
  param=$1
  RMI_PORT="com.sun.management.jmxremote.rmi.port"
  if [[ -n ${!param} ]]; then
    if [[ ! ${!param} == *"$RMI_PORT"* ]]; then
      echo "${param} should contain '$RMI_PORT' property. It is required for accessing the JMX metrics externally."
    fi
  fi
}

# Usage
# kafka_ready numBrokers timeout pathToConfig
function kafka_ready {
  if java $KAFKA_OPTS -cp "$CUB_CLASSPATH" "io.confluent.admin.kafka.health.KafkaReady"  $1 $2 $3
  then
    echo "Kafka ready: found at least $1 broker(s)."
  else
    exit 1
  fi
}

function log_status {
  echo "===> ${1}..."
}