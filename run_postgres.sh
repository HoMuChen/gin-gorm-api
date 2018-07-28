#!/bin/sh

docker run                        \
  -d                              \
  --name pgdb                     \
  -p $PG_PORT:5432                \
  -e POSTGRES_USER="$PG_USER"     \
  -e POSTGRES_PASSWORD="$PG_PWD"  \
  -v "$PWD/pgdata":/var/lib/postgresql/data \
  postgres
