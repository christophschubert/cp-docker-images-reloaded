# Scratch pad

Notes and tools used while developing the images.

```shell
# the following is a pattern to ensure that KAFKA_SUPER_PROP_2 gets set unless it has not been set before
#export KAFKA_SUPER_PROP_2
: "${KAFKA_SUPER_PROP_2:=byebye-world}"
```
Useful links:
- https://wiki.bash-hackers.org/syntax/pe#indirection