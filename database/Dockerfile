FROM postgres:14.5

COPY ./postgres-data /var/lib/postgresql/data
COPY ./init /docker-entrypoint-initdb.d/
