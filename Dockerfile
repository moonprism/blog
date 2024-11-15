FROM archlinux:latest

WORKDIR /app

COPY blog .
COPY dict dict

ENTRYPOINT ["./blog"]
