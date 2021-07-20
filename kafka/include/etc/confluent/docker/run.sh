#!/bin/bash

echo "Checking for required configuration settings..."


function exit_if_not_set {
  param=$1
  if [[ -z ${!param} ]]
  then
    echo "  Required environment variable $param not set"
    exit 1
  fi
}


# check whether we are running in KRAFT mode (this is done using the process.roles property)
if [[ -z "$KAFKA_PROCESS_ROLES" ]]
then
  echo "> configuring ZooKeeper mode"
  exit_if_not_set KAFKA_ZOOKEEPER_CONNECT
  LEGACY_MODE="true"
else
  echo "> configuring KRaft controller mode"
  # as the zookeeper setting will eventually go away, we take KRaft as default
  exit_if_not_set CLUSTER_ID
  exit_if_not_set KAFKA_NODE_ID
  exit_if_not_set KAFKA_CONTROLLER_QUORUM_VOTERS
fi


CONFIG_DIR=/etc/confluent/kafka
mkdir -p $CONFIG_DIR

SERVER_PROPERTIES_PATH=$CONFIG_DIR/server.properties


# TODO: ensure the LOG_DIRS is set -- is this necessary?
# TODO: KAFKA_DATA_DIR?


ub propertiesFromEnv /etc/confluent/docker/kafkaConfigSpec.json > $SERVER_PROPERTIES_PATH
ub formatLogger /etc/confluent/docker/log4j.properties.template /etc/confluent/docker/loggerDefaults.json KAFKA_LOG4J_ROOT_LOGLEVEL KAFKA_LOG4J_LOGGERS > /etc/kafka/log4j.properties

# TODO: format tools logger

if [[ -z $LEGACY_MODE ]]; then
  /usr/bin/kafka-storage format --ignore-formatted -t $CLUSTER_ID -c $SERVER_PROPERTIES_PATH
fi

/usr/bin/kafka-server-start $SERVER_PROPERTIES_PATH