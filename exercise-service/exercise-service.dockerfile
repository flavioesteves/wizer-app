FROM alpine:latest

RUN mkdir /app

COPY bin/exerciseApp /app

CMD ["app/exerciseApp"]
