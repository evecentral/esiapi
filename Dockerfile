FROM golang:1.10

WORKDIR /go/src/github.com/evecentral/esiapi
COPY . .
RUN curl https://glide.sh/get | sh
RUN glide install --strip-vendor
RUN go build ./cli/esi_stations_import/esi_stations_import.go
RUN go build ./cli/esi_token_tool

