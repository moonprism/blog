FROM mysql:5.7

ADD my.cnf /etc/mysql/
ADD *.sql /docker-entrypoint-initdb.d/

RUN chmod a+x /docker-entrypoint-initdb.d/*.sql

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
