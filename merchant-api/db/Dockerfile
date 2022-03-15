FROM mysql:latest

WORKDIR /docker-entrypoint-initdb.d

ADD ./sql/*.sql ./docker-entrypoint-initdb.d

EXPOSE 3306
