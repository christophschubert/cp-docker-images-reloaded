#!/bin/bash

set -e

echo "Checking for required configuration settings..."

# TODO: move to bash_functions.sh
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
# check for required properties
exit_if_not_set SCHEMA_REGISTRY_KAFKASTORE_BOOTSTRAP_SERVERS
exit_if_not_set SCHEMA_REGISTRY_HOST_NAME

# fail if any deprecated values are used
ub check-deprecated SCHEMA_REGISTRY_KAFKASTORE_CONNECTION_URL kafkastore.connection.url SCHEMA_REGISTRY_KAFKASTORE_BOOTSTRAP_SERVERS
ub check-deprecated SCHEMA_REGISTRY_PORT port SCHEMA_REGISTRY_LISTENERS

## TODO: discuss should we fail on other deprecated properties?
CONFIG_DIR=/etc/schema-registry
mkdir -p $CONFIG_DIR

PROPERTIES_PATH=$CONFIG_DIR/schema-registry.properties

ub propertiesFromEnvPrefix SCHEMA_REGISTRY > $PROPERTIES_PATH
ub propertiesFromEnvPrefix SCHEMA_REGISTRY_KAFKASTORE > $CONFIG_DIR/admin.properties
#ub formatLogger /etc/confluent/docker/log4j.properties.template /etc/confluent/docker/loggerDefaults.json KAFKA_LOG4J_ROOT_LOGLEVEL KAFKA_LOG4J_LOGGERS > /etc/kafka/log4j.properties

kafka_ready 1 20000 $SCHEMA_REGISTRY_KAFKASTORE_BOOTSTRAP_SERVERS $CONFIG_DIR/admin.properties
/usr/bin/schema-registry-start $PROPERTIES_PATH