FROM alpine:3.4

ADD beeso_linux /app
ADD plugins /app

EXPOSE 8080

CMD ["/app/beeso_linux"]