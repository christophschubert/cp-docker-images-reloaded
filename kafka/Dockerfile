FROM cp-base-reloaded

EXPOSE 9092

RUN rpm --import https://packages.confluent.io/rpm/6.2/archive.key \
    && printf "[Confluent.dist]\nname=Confluent repository (dist)\nbaseurl=https://packages.confluent.io/rpm/6.2/$releasever\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n\n[Confluent]\nname=Confluent repository\nbaseurl=https://packages.confluent.io/rpm/6.2\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n" > /etc/yum.repos.d/confluent.repo \
    && microdnf --nodocs install confluent-kafka-6.2.0 \
    && microdnf clean all \
    && rm -rf /tmp/*
# microdnf seems to lack an 'check-update' option -- TODO: is this critical?
COPY include/etc/confluent/docker /etc/confluent/docker

CMD /etc/confluent/docker/run.sh

