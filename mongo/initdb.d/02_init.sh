mongoimport -u superuser -p password123 --db sample --collection sample_collection --authenticationDatabase admin --file /docker-entrypoint-initdb.d/sample.json --jsonArray