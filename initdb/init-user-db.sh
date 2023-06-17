#!/bin/bash
set -e

POSTGRES_DB=k8s_go_opa  

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE $POSTGRES_DB;
EOSQL