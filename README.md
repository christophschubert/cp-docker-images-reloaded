# Confluent Docker images reloaded

Rewrite of the Confluent Platform Docker images.

Goals:
- Use of `microdnf` for all package management instead of `yum`
- Static Go binary `ub` instead of Python scripts to generate configurations, hence no runtime-dependency on Python (smaller image and smaller attack surface)
- Support for starting cluster in KRaft mode
- Deprecate all references to ZooKeeper for non-broker components  
- cleaned up configuration handling (maybe)

## Building
Currently, the `build.sh` script can be used to build base and Kafka node images.

The `examples` folder contains `docker compose` based examples using the new images.

## Configuration

Confluent Platform components ultimately require a Java properties file. During startup of the container, this config file will be generated from the environment variables passed to the container.
- easier to set defaults in ConfigSpec

## Implementation notes

All components run under a user `appuser`

### Directories

- `/etc/confluent/docker` - contains templates, configs, and scripts for configuring the applications on startup.
- `/etc/${COMPONENT}/secrets`
- `/var/lib/${COMPONENT}/data` 