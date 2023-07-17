FROM golang:1.18-alpine

RUN mkdir -p /app
COPY . /app/
WORKDIR /app
# COPY . ./
RUN go build -o /app/promgraf -v main.go

EXPOSE 2020

ENTRYPOINT ["/app/promgraf"]
