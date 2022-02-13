FROM alpine

WORKDIR /www

COPY linux_amd64_api .
RUN mkdir ./options
COPY options/.ini ./options

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ENTRYPOINT ["./linux_amd64_api", "web"]
