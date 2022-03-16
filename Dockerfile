FROM golang:1.18 as go-build
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o getir-go-assigment ./cmd

EXPOSE 8080
CMD /build/getir-go-assigment

