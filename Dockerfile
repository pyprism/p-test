FROM golang:1.16.6-alpine3.14

RUN apk update && apk add --no-cache git build-base inotify-tools curl

WORKDIR /app
COPY . /app

RUN go mod download

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air

