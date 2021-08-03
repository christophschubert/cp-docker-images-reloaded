FROM registry.access.redhat.com/ubi8/ubi-minimal:latest as java-builder

RUN rpm --install https://cdn.azul.com/zulu/bin/zulu-repo-1.0.0-1.noarch.rpm \
    && microdnf install --nodocs zulu-11  && microdnf clean all
ENV JAVA_HOME=/usr/lib/jvm/zulu-11
WORKDIR /build
COPY readyness-checks/ .
RUN ls
RUN ./gradlew clean installDist


FROM registry.access.redhat.com/ubi8/go-toolset as go-builder

WORKDIR /build
COPY ub/ub.go .
RUN go build -ldflags="-w -s" -o ub .


FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV LANG="C.UTF-8"
RUN rpm --install https://cdn.azul.com/zulu/bin/zulu-repo-1.0.0-1.noarch.rpm \
    && microdnf install --nodocs openssl \
                             wget \
                             nmap-ncat \
                             tar \
                             procps-ng \
                             krb5-workstation \
                             iputils \
                             hostname  \
    && microdnf install --nodocs zulu-11 \
    && microdnf clean all \
    && rm -rf /tmp/* \
    && mkdir -p /etc/confluent/docker /usr/logs \
    && useradd --no-log-init --create-home --shell /bin/bash appuser \
    && chown appuser:appuser -R /etc/confluent /usr/logs

# microdnf seems to lack an 'check-update' option -- TODO: is this critical?

ENV CUB_CLASSPATH=/usr/share/java/cp-base-new/*
# avoid having to build cp-base package (contains 'utility belt' functions)
COPY include/etc/confluent/docker /etc/confluent/docker

# selectively copy necessary jars (without copying, e.g., compression libraries)
COPY --from=java-builder /build/kafka-ready/build/install/kafka-ready/lib/kafka*.jar /usr/share/java/cp-base-new/
COPY --from=java-builder /build/kafka-ready/build/install/kafka-ready/lib/slf4j*.jar /usr/share/java/cp-base-new/
COPY --from=java-builder /build/kafka-ready/build/install/kafka-ready/lib/log4j*.jar /usr/share/java/cp-base-new/

COPY --from=go-builder /build/ub /usr/bin
