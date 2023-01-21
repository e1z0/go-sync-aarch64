#!/bin/sh -e

if [ ! -e "/db/shared-local-instance.db" ]; then
echo "Database does not exist, creating one..."
cp /db-orig/* /db/
fi

java -jar DynamoDBLocal.jar -sharedDb -dbPath /db
