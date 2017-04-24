FROM alpine

ADD beeso_linux /
ADD plugins /plugins

EXPOSE 8080

CMD /beeso_linux