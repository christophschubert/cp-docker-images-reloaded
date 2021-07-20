# Scratch pad



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

Notes and tools used while developing the images.

```shell
# the following is a pattern to ensure that KAFKA_SUPER_PROP_2 gets set unless it has not been set before
#export KAFKA_SUPER_PROP_2
: "${KAFKA_SUPER_PROP_2:=byebye-world}"
```
Useful links:
- https://wiki.bash-hackers.org/syntax/pe#indirection



Size comparison:
cp-kafka-reloaded                             latest                                                  d756d798f8a1   12 minutes ago   453MB (without toole)
cp-kafka-reloaded                             latest                                                  c04f71aca8dd   7 minutes ago    516MB (with tools)
cp-kafka-reloaded                             latest                                                  4c5c36e6bc9b   6 seconds ago    508MB (with tools and --nodocs option)
cp-base-reloaded                              latest                                                  982c76bd9ed6   2 hours ago      106MB (without tools)
confluentinc/cp-enterprise-kafka              6.2.0                                                   a6215cd2f450   2 weeks ago      901MB
confluentinc/cp-kafka                         6.2.0                                                   ca0dbcd0244c   2 weeks ago      771MB
confluentinc/cp-server                        6.2.0                                                   81958e2b59e0   2 weeks ago      1.44GB
confluentinc/cp-schema-registry               6.2.0                                                   e560c8baba46   2 weeks ago      1.52GB
confluentinc/cp-zookeeper                     6.2.0                                                   04999d93068f   2 weeks ago      771MB
confluentinc/cp-enterprise-replicator         6.2.0                                                   adb9781e9407   6 weeks ago      2.14GB
confluentinc/cp-server-connect                6.0.1                                                   90851c79a2b1   8 months ago     1.65GB
confluentinc/cp-enterprise-kafka              6.0.1                                                   fa9b92eac717   8 months ago     820MB
confluentinc/cp-kafka                         6.0.1                                                   64f8db8ddbe5   8 months ago     714MB
confluentinc/cp-zookeeper                     6.0.1                                                   0667a7f01cf0   8 months ago     714MB
confluentinc/cp-schema-registry               6.0.1                                                   c7dfd2529fe3   8 months ago     1.34GB
confluentinc/cp-enterprise-kafka              5.3.1                                                   2b5ee743a097   22 months ago    646MB