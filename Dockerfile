FROM archlinux:latest

WORKDIR /app

COPY main .
COPY dict dict

ENTRYPOINT ["./main"]
