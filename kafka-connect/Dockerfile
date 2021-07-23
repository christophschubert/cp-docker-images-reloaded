FROM cp-kafka-reloaded

ENV CONFLUENT_VERSION=6.2.0
ENV COMPONENT=kafka-connect

EXPOSE 8083
ARG FILL_ME
ARG HUB_COMPONENTS_PATH=/usr/share/confluent-hub-components

USER root
#mind the extra '/' below:
ENV CONNECT_PLUGIN_PATH=/usr/share/java/,$HUB_COMPONENTS_PATH/

RUN echo "===> Installing components ..." \
  && microdnf install --nodocs \
    confluent-schema-registry-${CONFLUENT_VERSION} \
    confluent-control-center-${CONFLUENT_VERSION} \
    confluent-hub-client-${CONFLUENT_VERSION} \
  && echo "===> Cleaning up ..." \
  && microdnf clean all \
  && rm -rf /tmp/* \
  && rm -rf /etc/confluent-control-center /usr/bin/control-center-* /usr/share/doc/confluent-control-center /usr/share/java/confluent-control-center \
  && rm -rf /etc/schema-registry /usr/bin/schema-registry* /usr/share/doc/kafka-serde-tools /usr/share/doc/schema-registry /usr/share/java/schema-registry \
  && echo "===> Setting up ${COMPONENT} dirs ..." \
  && mkdir -p /etc/${COMPONENT} /etc/${COMPONENT}/secrets/ /etc/${COMPONENT}/jars \
  && chmod -R 700 /etc/${COMPONENT} \
  && chown -R appuser:appuser /etc/$COMPONENT \
  && mkdir -p ${HUB_COMPONENTS_PATH} \
  && chown appuser:appuser -R ${HUB_COMPONENTS_PATH}



# TODO: discuss proper access for /etc/kafka-connect, /etc/kafka-connect/jars, /etc/kafka-connect/secrets
# TODO: discuss whether we should remove /etc/kafka from this image (not shrink size, but might cause less confusion)
# TODO: why do we use /etc/kafka-connect/jars here? If this is a valid mount point for a volume?
# See also: https://medium.com/@nielssj/docker-volumes-and-file-system-permissions-772c1aee23ca
#  In this case, it might get mounted to /usr/share?
VOLUME ["/etc/${COMPONENT}/jars", "/etc/${COMPONENT}/secrets"]

COPY --chown=appuser:appuser include/etc/confluent/docker /etc/confluent/docker

USER appuser

CMD ["/etc/confluent/docker/run"]

#TODO: add healthcheck