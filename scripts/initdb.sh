#!/bin/bash
psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f /dump.sql -v ON_ERROR_STOP=1 -f /dump.sql
