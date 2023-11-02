#####################################
#   STEP 1: build golang executable #
#####################################
FROM golang:1.20-alpine3.18 AS build

#ENV GO111MODULE on
RUN mkdir -p /etc/secret/
COPY  .env /etc/secret/.env

ARG ENV
ENV ENV $ENV

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN CGO_ENABLED=0 go build -a -o main

#####################################
#   STEP 2 build liquibase image    #
#####################################
FROM liquibase/liquibase:alpine

USER root
ARG ENV
ENV ENV $ENV

COPY ./config/changelog.xml /liquibase/changelog/changelog.xml
ADD ./sql /liquibase/scripts

ENV LIQUIBASE_CHANGELOG_FILE changelog.xml

COPY --from=build /go/release/main /app/main

WORKDIR /app/
RUN chmod +x /app/main
ENTRYPOINT ["./main"]