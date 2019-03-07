#!/bin/bash

docker run --name baradb -p 5432:5432 -e POSTGRES_PASSWORD=root -d postgres:9.2-alpine
