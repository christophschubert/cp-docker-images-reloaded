#!/usr/bin/env bash

# include helper functions
. /etc/confluent/docker/bash-functions.sh


log_status "Checking for required configuration settings"
# check whether we are running in KRAFT mode (this is done using the process.roles property)
if [[ -z "$KAFKA_PROCESS_ROLES" ]]; then
  echo "> configuring ZooKeeper mode"
  LEGACY_MODE="true"
  exit_if_not_set KAFKA_ZOOKEEPER_CONNECT
  exit_if_not_set KAFKA_ADVERTISED_LISTENERS
else
  echo "> configuring KRaft controller mode"
  # as the zookeeper setting will eventually go away, we take KRaft as default
  exit_if_not_set CLUSTER_ID
  exit_if_not_set KAFKA_NODE_ID
  exit_if_not_set KAFKA_CONTROLLER_QUORUM_VOTERS
  if [[ $KAFKA_PROCESS_ROLES == *"broker"* ]]; then
    exit_if_not_set KAFKA_ADVERTISED_LISTENERS
  fi
fi
# TODO: seems a single broker/single controller setup can't start when using this defaults: are these sensible?
if [ -z $KAFKA_LISTENERS ]; then
  KAFKA_LISTENERS=$(ub listeners "$KAFKA_ADVERTISED_LISTENERS")
  export KAFKA_LISTENERS
fi
ub check-deprecated KAFKA_ADVERTISED_HOST advertised.host KAFKA_ADVERTISED_LISTENERS
ub check-deprecated KAFKA_ADVERTISED_PORT advertised.port KAFKA_ADVERTISED_LISTENERS
ub check-deprecated KAFKA_HOST host KAFKA_ADVERTISED_LISTENERS
ub check-deprecated KAFKA_PORT port KAFKA_ADVERTISED_LISTENERS


# TODO: ensure the LOG_DIRS is set -- is this necessary?
# TODO: KAFKA_DATA_DIR?

CONFIG_DIR=/etc/confluent/kafka
mkdir -p $CONFIG_DIR

SERVER_PROPERTIES_PATH=$CONFIG_DIR/server.properties

# TODO: why do we use two different paths here?
ub propertiesFromEnv /etc/confluent/docker/kafkaConfigSpec.json > $SERVER_PROPERTIES_PATH
ub formatLogger /etc/confluent/docker/log4j.properties.template /etc/confluent/docker/loggerDefaults.json KAFKA_LOG4J_ROOT_LOGLEVEL KAFKA_LOG4J_LOGGERS > /etc/kafka/log4j.properties

# TODO: format tools logger


if [[ -z $LEGACY_MODE ]]; then
  /usr/bin/kafka-storage format --ignore-formatted -t $CLUSTER_ID -c $SERVER_PROPERTIES_PATH
fi
log_status "Launching Kafka node"
/usr/bin/kafka-server-start $SERVER_PROPERTIES_PATH