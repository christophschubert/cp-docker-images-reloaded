#!/usr/bin/env bash

. /etc/confluent/docker/bash-functions.sh

exit_if_not_set CONNECT_BOOTSTRAP_SERVERS
exit_if_not_set CONNECT_GROUP_ID
exit_if_not_set CONNECT_CONFIG_STORAGE_TOPIC
exit_if_not_set CONNECT_OFFSET_STORAGE_TOPIC
exit_if_not_set CONNECT_STATUS_STORAGE_TOPIC

# TODO: why do we need to explicitly set key and value converter?

exit_if_not_set CONNECT_ADVERTISED_HOST_NAME

PROPERTIES_PATH=/etc/"${COMPONENT}"/kafka-connect.properties
ub propertiesFromEnvPrefix CONNECT > $PROPERTIES_PATH

#TODO: configure Log4J

echo "===> check Kafka is healthy"
# TODO: wouldn't is make more sense to check the configured replication factors instead of hardcoding 1?
kafka_ready 1 40000 $CONNECT_BOOTSTRAP_SERVERS $PROPERTIES_PATH

# TODO: setup JMX

if [ -z "$CLASSPATH" ]; then
  export CLASSPATH="/etc/kafka-connect/jars/*"
fi
connect-distributed $PROPERTIES_PATH