FROM golang:latest AS gobuild
WORKDIR /build-dir
COPY go.mod .
RUN go mod download all
COPY . .
RUN go build -o /tmp/ test_task. 

FROM debian AS app 
RUN apt-get update
COPY --from=gobuild /tmp/test_task /app/test_task
EXPOSE 8080
CMD ["/app/test_task"]