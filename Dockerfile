FROM reg.yunpro.cn/dingqiwei/alpine

ADD . /app

WORKDIR /app

EXPOSE 8080

CMD ["/app/beeso"]