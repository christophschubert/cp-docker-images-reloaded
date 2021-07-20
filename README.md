# Confluent Docker images reloaded

Rewrite of the Confluent Platform Docker images.

Goals:
- Use of `microdnf` for all package management instead of `yum`
- Static Go binary `ub` instead of Python scripts to generate configurations, hence no runtime-dependency on Python (smaller image and smaller attack surface)
- Support for starting cluster in KRaft mode
- cleaned up configuration handling (maybe)

## Building
Currently, the `build.sh` script can be used to build base and Kafka node images.

The `examples` folder contains `docker compose` based examples using the new images. 