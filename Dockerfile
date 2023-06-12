FROM golang:latest

RUN apt update && \
	apt full-upgrade -y

WORKDIR /usr/api/

COPY . .

RUN git config --global --add safe.directory /usr/api
RUN go mod tidy && \
	go install github.com/swaggo/swag/cmd/swag@latest

CMD ["tail", "-f", "/dev/null"]
