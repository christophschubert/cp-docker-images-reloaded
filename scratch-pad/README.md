# Scratch pad

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