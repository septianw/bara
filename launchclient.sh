#!/bin/bash
docker run -it --rm --link baradb:postgres postgres:9.2-alpine psql -h postgres -U postgres
