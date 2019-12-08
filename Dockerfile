FROM golang:1.8-alpine
COPY . /go/src/MyApp
WORKDIR /go/src/MyApp
RUN go install MyApp

FROM alpine:latest
COPY --from=0 /go/bin/MyApp .
ENV PORT 6969
EXPOSE 6969
ENTRYPOINT ["./MyApp"]
