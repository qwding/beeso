FROM alpine:3.4

ADD beeso_linux /
ADD plugins /plugins

EXPOSE 8080

CMD /beeso_linux