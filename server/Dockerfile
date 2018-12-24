FROM golang:1.11
ENV GO111MODULE=on
WORKDIR /api

COPY . .
RUN go get github.com/pilu/fresh

CMD ["fresh"]
