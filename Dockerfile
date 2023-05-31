FROM alpine:latest
WORKDIR /workdir
COPY ./building /app
ENTRYPOINT /app