FROM golang:1.23

WORKDIR /mnemstart
COPY . /mnemstart/
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -o main .
EXPOSE 3000
ENTRYPOINT ["./main"]
