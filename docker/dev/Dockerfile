## Build
FROM golang:alpine3.18

USER $USER_ID:$USER_ID

WORKDIR /app

COPY start.sh /

ENTRYPOINT /start.sh

CMD ["/bin/ash"]