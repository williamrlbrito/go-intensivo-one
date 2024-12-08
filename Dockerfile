FROM golang:1.21.4
WORKDIR /app
ENTRYPOINT [ "tail", "-f", "/dev/null" ]