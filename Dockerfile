FROM --platform=linux/amd64 golang:1.19-buster

WORKDIR /app

ADD . /app
RUN env GOOS=linux go build -o .

RUN apt-get update && apt-get install -y --allow-unauthenticated tzdata

EXPOSE 1323

ENTRYPOINT ["./gcp-cloud-run-boilerplate-golang"]
