FROM alpine

WORKDIR /www

COPY linux_amd64_api .

CMD mkdir data
COPY data/*.txt ./data/

COPY options/.ini ./options/

ENTRYPOINT ["./linux_amd64_api", "web", "--reindex"]
