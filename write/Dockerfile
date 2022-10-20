FROM alpine

WORKDIR /www

COPY linux_amd64_api .

RUN mkdir ./data
COPY data/*.text ./data/

COPY options/.ini ./options/

ENTRYPOINT ["./linux_amd64_api", "web"]
