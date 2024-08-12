FROM alpine:latest

RUN mkdir /app

COPY migrations /app/migrations

COPY bin/migrationApp /app


ENTRYPOINT ["app/migrationApp", "migrate","up"]
