#!/bin/bash

/usr/bin/kafka-storage format -t 1_TU6DOhT9uVDTk3B3E8Zg -c /etc/kafka/kraft/server.properties
/usr/bin/kafka-server-start /etc/kafka/kraft/server.properties