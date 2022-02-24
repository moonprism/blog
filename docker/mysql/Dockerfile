FROM mysql:5.7

ADD *.sql /docker-entrypoint-initdb.d/
RUN chmod a+x /docker-entrypoint-initdb.d/*.sql
