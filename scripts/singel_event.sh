#!/usr/bin/env bash

set -xe

ID=${ID:-"1"}
curl "http://localhost:8080/events" \
     -H "Content-Type: application/json" \
     -d "{\"id\": \"${ID}\", \"someIntVal\": 42, \"someFloatVal\": 3.14}"
