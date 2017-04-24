FROM alpine:3.4

ADD beeso /app
ADD plugins /app/plugins

EXPOSE 8080

CMD ["/app/beeso"]