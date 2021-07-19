#!/bin/bash

echo "Checking for required configuration settings..."

#see https://wiki.bash-hackers.org/syntax/pe#indirection for the origin of this
function exit_if_not_set {
  param=$1
  if [[ -z ${!param} ]]
  then
    echo "  Required environment variable $param not set"
    exit 1
  fi
}


# check whether we are running in KRAFT mode (this is done using the process.roles property)
if [[ -n "$KAFKA_PROCESS_ROLES" ]]
then
  echo "> configuring ZooKeeper mode"
  # set KRaft-mode = false
else
  echo "> configuring KRaft controller mode"
  # as the zookeeper setting will eventually go away, we take KRaft as default
  exit_if_not_set KAFKA_NODE_ID
#  exit_if_not_set
  #check for node id, controller connection string
fi

CONFIG_DIR=/etc/confluent/kafka
mkdir -p $CONFIG_DIR

SERVER_PROPERTIES_PATH=$CONFIG_DIR/server.properties

# TODO: ensure the LOG_DIRS is set -- is this necessary?
ub envToProp KAFKA > $SERVER_PROPERTIES_PATH
ub envToProp CONFLUENT >> $SERVER_PROPERTIES_PATH #TODO: change to keep the 'confluent.' prefix

#TODO: get cluster-id from config.
#TODO: ensure kafka-storage gets called only when we are in KRaft-mode
/usr/bin/kafka-storage format --ignore-formatted -t 1_TU6DOhT9uVDTk3B3E8Zg -c $SERVER_PROPERTIES_PATH
/usr/bin/kafka-server-start $SERVER_PROPERTIES_PATH