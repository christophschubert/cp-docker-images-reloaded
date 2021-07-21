#!/bin/bash

# include helper functions
. /etc/confluent/docker/bash-functions.sh


echo "Checking for required configuration settings..."
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