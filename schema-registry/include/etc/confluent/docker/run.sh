#!/bin/bash

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



exit_if_not_set SCHEMA_REGISTRY_HOST_NAME


CONFIG_DIR=/etc/schema-registry
mkdir -p $CONFIG_DIR

PROPERTIES_PATH=$CONFIG_DIR/schema-registry.properties



ub propertiesFromEnv /etc/confluent/docker/schemaRegistryConfigSpec.json > $PROPERTIES_PATH
#ub formatLogger /etc/confluent/docker/log4j.properties.template /etc/confluent/docker/loggerDefaults.json KAFKA_LOG4J_ROOT_LOGLEVEL KAFKA_LOG4J_LOGGERS > /etc/kafka/log4j.properties


/usr/bin/schema-registry-start $PROPERTIES_PATH