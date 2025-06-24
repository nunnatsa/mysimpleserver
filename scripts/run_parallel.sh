#!/usr/bin/env bash

set -xe

for i in $(seq 1 50); do
  for j in $(seq 1 100); do
    curl "http://localhost:8080/events" \
         -H "Content-Type: application/json" \
         -d "{\"id\": \"$i\", \"someIntVal\": 42, \"someFloatVal\": 3.14}" &
  done
done