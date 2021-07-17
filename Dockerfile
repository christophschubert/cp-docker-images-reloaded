FROM registry.access.redhat.com/ubi8/ubi-minimal:8.4-205

RUN rpm --import https://packages.confluent.io/rpm/6.2/archive.key && \
    printf "[Confluent.dist]\nname=Confluent repository (dist)\nbaseurl=https://packages.confluent.io/rpm/6.2/$releasever\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n\n[Confluent]\nname=Confluent repository\nbaseurl=https://packages.confluent.io/rpm/6.2\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n" > /etc/yum.repos.d/confluent.repo && \
    rpm --install https://cdn.azul.com/zulu/bin/zulu-repo-1.0.0-1.noarch.rpm && \
    microdnf install zulu11-jdk-headless --nodocs && \
    microdnf install confluent-kafka-6.2.0 && \
    microdnf clean all && \
    rm -rf /tmp/*

