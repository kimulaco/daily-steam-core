#!/bin/bash

echo "./bin/cloud_sql_proxy -instances=daily-steam:us-central1:daily-steam-db=tcp:$DB_PORT"
./bin/cloud_sql_proxy -instances=daily-steam:us-central1:daily-steam-db=tcp:$DB_PORT
