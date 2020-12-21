FROM golang:1.10

WORKDIR /login
COPY . /login

EXPOSE 8000

CMD ["go", "run", "main.go"]