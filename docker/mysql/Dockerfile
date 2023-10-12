FROM mysql:8.1

ADD *.sql /docker-entrypoint-initdb.d/
RUN chmod a+x /docker-entrypoint-initdb.d/*.sql
