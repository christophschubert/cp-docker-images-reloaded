FROM cp-base-reloaded

EXPOSE 9092

USER root

ENV COMPONENT=kafka
ENV CONFLUENT_VERSION=6.2.0
RUN rpm --import https://packages.confluent.io/rpm/6.2/archive.key \
    && printf "[Confluent.dist]\nname=Confluent repository (dist)\nbaseurl=https://packages.confluent.io/rpm/6.2/$releasever\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n\n[Confluent]\nname=Confluent repository\nbaseurl=https://packages.confluent.io/rpm/6.2\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n" > /etc/yum.repos.d/confluent.repo \
    && microdnf --nodocs install confluent-server-${CONFLUENT_VERSION} \
    && microdnf --nodocs install confluent-rebalancer-${CONFLUENT_VERSION} \
    && microdnf --nodocs install confluent-security-${CONFLUENT_VERSION} \
    # TODO: double check why kafka-rest is installed
    && echo "===> Clean up" \
    && microdnf clean all \
    && rm -rf /tmp/* \
    && rm -rf /usr/share/java/confluent-security/schema-registry /usr/share/java/confluent-security/ksql /usr/share/java/confluent-security/kafka-rest /usr/share/java/confluent-security/connect \
    && echo "===> Setting up ${COMPONENT} dirs" \
    && mkdir -p /var/lib/${COMPONENT}/data /etc/${COMPONENT}/secrets \
    && chmod -R ag+w /etc/kafka /var/lib/${COMPONENT}/data /etc/${COMPONENT}/secrets \
    && chown -R appuser:appuser /var/log/kafka /var/log/confluent /var/lib/kafka /var/lib/zookeeper /etc/${COMPONENT}/secrets

COPY --chown=appuser:appuser include/etc/confluent/docker /etc/confluent/docker

USER appuser

CMD /etc/confluent/docker/run.sh