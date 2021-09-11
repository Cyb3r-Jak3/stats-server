FROM golang:1.17-alpine as builder

WORKDIR /src

WORKDIR /go/src/app
COPY . ./

ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go build -o /go/bin/app

FROM gcr.io/distroless/static
COPY --from=builder /go/bin/app /
ENV PRODUCTION=TRUE
EXPOSE 5000
CMD ["/app"]