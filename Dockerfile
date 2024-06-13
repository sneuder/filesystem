FROM golang:1.21.1 AS dev

WORKDIR /app

RUN apt update -y
RUN apt upgrade -y

RUN apt install make -y

CMD ["sleep", "infinity"]

