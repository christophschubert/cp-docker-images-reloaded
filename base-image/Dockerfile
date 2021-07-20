FROM registry.access.redhat.com/ubi8/go-toolset as builder

WORKDIR /build
COPY ub/ub.go .
RUN go build -ldflags="-w -s" -o ub .

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
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
    && mkdir -p /etc/confluent/docker

# microdnf seems to lack an 'check-update' option -- TODO: is this critical?

COPY --from=builder /build/ub /usr/bin
