FROM debian:9

RUN mkdir /www
RUN mkdir /www/static
RUN mkdir /www/config
RUN mkdir /www/web

COPY ./write /www
RUN chmod +x /www/write
COPY ./prod.ini /www/config
COPY ./dist /www/web

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /www

ENTRYPOINT ["./write", "--env", "test"]