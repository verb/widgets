FROM golang:1.14-alpine as build

WORKDIR /go/src/app
ADD main.go /go/src/app
ENV CGO_ENABLED=0
RUN go build -o /go/bin/app

FROM gcr.io/distroless/static-debian10
COPY --from=build /go/bin/app /
CMD ["/app"]
