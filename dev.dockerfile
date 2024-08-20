FROM golang:1.23

WORKDIR /mnemstart
COPY . /mnemstart/
RUN go mod download
RUN go mod verify
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/air-verse/air@latest
CMD air -c .air.toml
