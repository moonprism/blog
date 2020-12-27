FROM debian:9

RUN mkdir -p /www/static /www/log
COPY . /www
RUN chmod +x /www/main

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /www

ENTRYPOINT ["./main", "--env", "prod", "--reindex"]
