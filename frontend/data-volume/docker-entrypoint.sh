#!/bin/bash
set -e
chown -R postgres:postgres /var/local/event-collector.csv
chmod 777 /var/local/event-collector.csv
exec "$@"