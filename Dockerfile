# ------------------------------------------------------
FROM golang:1.11-alpine as builder

LABEL "name"="gh-stars"
LABEL "maintainer"="Augusto Pascutti <augusto.hp+github@gmail.com>"
LABEL "version"="1.1.0"

LABEL "com.github.actions.name"="GitHub Stars"
LABEL "com.github.actions.descriptions"="Lists your starred items from past week"
LABEL "com.github.actions.icon"="star"
LABEL "com.github.actions.color"="yellow"

RUN apk add --no-cache git make

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN make build
RUN cp "${GOPATH}/bin/gh-stars" "/usr/bin/gh-stars"

# ------------------------------------------------------
FROM alpine:latest as app

COPY --from=builder /usr/bin/gh-stars /usr/bin/gh-stars

ENTRYPOINT ["/usr/bin/gh-stars"]

CMD [""]
