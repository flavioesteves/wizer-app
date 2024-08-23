FROM alpine:latest

RUN mkdir /app

COPY bin/routineApp /app

CMD ["app/routineApp"]
