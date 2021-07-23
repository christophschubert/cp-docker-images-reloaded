#!/usr/bin/env bash

# rudimentary test script for include/etc/confluent/docker/bash-functions.sh

. include/etc/confluent/docker/bash-functions.sh

EXISTING_ENV_VAR=world
exit_if_not_set EXISTING_ENV_VAR # test failure would quit whole test script

warn_jmx_rmi_port EXISTING_ENV_VAR # should output warning
PROPER_JMX=com.sun.management.jmxremote.rmi.port=sdfkdfd
warn_jmx_rmi_port PROPER_JMX # expect no output
warn_jmx_rmi_port NON_EXISTING_ENV_VAR # no output

(exit_if_not_set NON_EXISTING_ENV_VAR ) && exit 1 # to check that exit_if_not_set fails

echo "all passed"


