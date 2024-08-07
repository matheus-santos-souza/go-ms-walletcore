<<<<<<< HEAD
FROM golang:1.20
=======
FROM golang:1.22
>>>>>>> master

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]