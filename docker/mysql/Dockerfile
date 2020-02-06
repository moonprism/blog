FROM mysql:5.7

COPY ./my.cnf /etc/mysql/my.cnf
COPY ./blog.sql /docker-entrypoint-initdb.d/dump.sql

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone