

Base image: 

Steps run from inside the base image:
```shell
rpm --import https://packages.confluent.io/rpm/6.2/archive.key
```



```shell
printf "[Confluent.dist]\nname=Confluent repository (dist)\nbaseurl=https://packages.confluent.io/rpm/6.2/$releasever\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n\n[Confluent]\nname=Confluent repository\nbaseurl=https://packages.confluent.io/rpm/6.2\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n"
printf "[Confluent.dist]\nname=Confluent repository (dist)\nbaseurl=https://packages.confluent.io/rpm/6.2/$releasever\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n\n[Confluent]\nname=Confluent repository\nbaseurl=https://packages.confluent.io/rpm/6.2\ngpgcheck=1\ngpgkey=https://packages.confluent.io/rpm/6.2/archive.key\nenabled=1\n" > /etc/yum.repos.d/confluent.repo
```
(can use `microdnf repolist`) to check

microdnf install java-1.8.0-openjdk-headless --nodocs
microdnf install confluent-kafka-6.2.0


## Learnings:

Options for `kafka-storage`:
```shell
[root@672fb0ec2f46 /]# kafka-storage --help
usage: kafka-storage [-h] {info,format,random-uuid} ...

The Kafka storage tool.

positional arguments:
  {info,format,random-uuid}
    info                 Get information about the Kafka log directories on this node.
    format               Format the Kafka log directories on this node.
    random-uuid          Print a random UUID.

optional arguments:
  -h, --help             show this help message and exit
[root@672fb0ec2f46 /]# kafka-storage info --help
usage: kafka-storage info [-h] --config CONFIG

optional arguments:
  -h, --help             show this help message and exit
  --config CONFIG, -c CONFIG
                         The Kafka configuration file to use.
[root@672fb0ec2f46 /]# kafka-storage random-uuid --help
usage: kafka-storage random-uuid [-h]

optional arguments:
  -h, --help             show this help message and exit
[root@672fb0ec2f46 /]# kafka-storage format --help
usage: kafka-storage format [-h] --config CONFIG --cluster-id CLUSTER_ID [--ignore-formatted]

optional arguments:
  -h, --help             show this help message and exit
  --config CONFIG, -c CONFIG
                         The Kafka configuration file to use.
  --cluster-id CLUSTER_ID, -t CLUSTER_ID
                         The cluster ID to use.
  --ignore-formatted, -g
```