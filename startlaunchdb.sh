#!/bin/bash

docker run --name baradb -e POSTGRES_PASSWORD=root -d postgres:9.2-alpine
docker run -it --rm --link baradb:postgres postgres:9.2-alpine psql -h postgres -U postgres
