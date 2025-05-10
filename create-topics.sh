#!/bin/bash

# while ! nc -z kafka 9092; do
#   echo "Waiting for Kafka to be ready..."
#   sleep 2
# done


kafka-topics.sh --create --topic delete_card --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1 --if-not-exists
