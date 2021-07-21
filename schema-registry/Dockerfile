FROM cp-base-reloaded

ENV CONFLUENT_VERSION=6.2.0
ENV COMPONENT=schema-registry

RUN rpm --import https://packages.confluent.io/rpm/6.2/archive.key \
    && printf "[Confluent.dist]\nname=Confluent repository (dist)\nbaseurl=https://packages.confluent.io/rpm/6.2/$releasever\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n\n[Confluent]\nname=Confluent repository\nbaseurl=https://packages.confluent.io/rpm/6.2\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n" > /etc/yum.repos.d/confluent.repo \
    && microdnf --nodocs install confluent-${COMPONENT}-${CONFLUENT_VERSION} \
            confluent-control-center-${CONFLUENT_VERSION} \
            # We are installing confluent-telemetry package explicitly because
            # Schema Registry's deb/rpm packages cannot directly depend on
            # confluent-telemetry package as Schema Registry is Open Source.
            confluent-telemetry-${CONFLUENT_VERSION} \
            confluent-security-${CONFLUENT_VERSION} \
    && microdnf clean all \
    && rm -rf /tmp/* \
    && rm -rf /usr/share/java/confluent-control-center /usr/share/doc/confluent-control-center /usr/bin/control-center-* /etc/confluent-control-center \
    && rm -rf /usr/share/java/confluent-security/schema-validator /usr/share/java/confluent-security/ksql /usr/share/java/confluent-security/kafka-rest /usr/share/java/confluent-security/connect
# remove most of the contents of confluent-control-center (basically just keep monitoring interceptors) and the security components for other platforms
# TODO: double check whether this has any negative impact.
COPY include/etc/confluent/docker /etc/confluent/docker

CMD /etc/confluent/docker/run.sh