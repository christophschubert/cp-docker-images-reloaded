FROM cp-base-reloaded

ENV CONFLUENT_VERSION=6.2.0
ENV COMPONENT=schema-registry

USER root

RUN rpm --import https://packages.confluent.io/rpm/6.2/archive.key \
    && printf "[Confluent.dist]\nname=Confluent repository (dist)\nbaseurl=https://packages.confluent.io/rpm/6.2/$releasever\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n\n[Confluent]\nname=Confluent repository\nbaseurl=https://packages.confluent.io/rpm/6.2\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n" > /etc/yum.repos.d/confluent.repo \
    && microdnf --nodocs install confluent-${COMPONENT}-${CONFLUENT_VERSION} \
            confluent-control-center-${CONFLUENT_VERSION} \
            # We are installing confluent-telemetry package explicitly because
            # Schema Registry's deb/rpm packages cannot directly depend on
            # confluent-telemetry package as Schema Registry is Open Source.
            confluent-telemetry-${CONFLUENT_VERSION} \
            confluent-security-${CONFLUENT_VERSION} \
    && echo "===> clean up ...." \
    && microdnf clean all \
    && rm -rf /tmp/* \
    && rm -rf /usr/share/java/confluent-control-center /usr/share/doc/confluent-control-center /usr/bin/control-center-* /etc/confluent-control-center \
    && rm -rf /usr/share/java/confluent-security/schema-validator /usr/share/java/confluent-security/ksql /usr/share/java/confluent-security/kafka-rest /usr/share/java/confluent-security/connect \
    && echo "===> setting up directories ...." \
    && mkdir -p /etc/${COMPONENT}/secrets \
    && chown -R appuser:appuser /etc/${COMPONENT} \
    && chmod -R go-wxr /etc/${COMPONENT} /etc/${COMPONENT}/secrets
# remove most of the contents of confluent-control-center (basically just keep monitoring interceptors) and the security components for other platforms
# TODO: double check whether this has any negative impact.
# TODO: original image creates /usr/logs here, is this needed?
# TODO: changed access rights on /etc/schema-registry to be more restrictive
COPY --chown=appuser:appuser include/etc/confluent/docker /etc/confluent/docker

VOLUME ["/etc/${COMPONENT}/secrets"]

USER appuser

CMD /etc/confluent/docker/run.sh