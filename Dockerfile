FROM golang:latest

RUN apt update && \
	apt full-upgrade -y

WORKDIR /usr/api/

RUN git config --global --add safe.directory /usr/api

CMD ["tail", "-f", "/dev/null"]

