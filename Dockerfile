FROM ubuntu:16.04

ADD . /app

WORKDIR /app

EXPOSE 8080

CMD ["/app/beeso_linux"]