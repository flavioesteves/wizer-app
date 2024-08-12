FROM alpine:latest

RUN mkdir /app

COPY bin/profileApp /app

CMD ["app/profileApp"]
