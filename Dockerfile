  
# Dockerfile
FROM golang:latest AS builder
ENV GO111MODULE=on
ENV ADDRESS_DATA=https://gist.githubusercontent.com/hadeshunter/42f43e63d074ecc7a25ff0df5a90ad19/raw/b797fedd62288ccbe1c6b81bf7f6dc14703b8535/vi-address-2016.csv

# Download modules
WORKDIR $GOPATH/src/github.com/hadeshunter/todo
RUN mkdir /data && curl ${ADDRESS_DATA} -o /data/address.csv
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /todo .

# Run
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /todo ./
COPY --from=builder /data ./
EXPOSE 5000
ENTRYPOINT ["./todo"]