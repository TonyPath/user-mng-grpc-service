#!/bin/sh
set -e

migrate_db() {
    i=0
    ./migrate -database postgres://${PG_USERNAME}:${PG_PASSWORD}@${PG_HOST}/${PG_DBNAME}?sslmode=disable -verbose -source file://./migrations up
    while [ $? -ne 0 -a $i -lt 10 ]; do
        sleep 3
        i=`expr $i + 1`
        ./migrate -database postgres://${PG_USERNAME}:${PG_PASSWORD}@${PG_HOST}/${PG_DBNAME}?sslmode=disable -verbose -source file://./migrations up
    done
}

check_migration_success() {
    if [ $? -ne 0 ]; then
        exit 1
    fi
}

set +e

migrate_db
check_migration_success
echo "Migration completed"

./user-mng-service
