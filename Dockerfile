FROM golang:1.25.4
RUN mkdir /app
WORKDIR /app
COPY cmd/ /app/cmd/
COPY pkg/ /app/pkg/
COPY go.mod /app/
RUN go mod download
RUN go mod tidy
