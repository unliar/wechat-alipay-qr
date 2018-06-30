FROM ubuntu:18.04
WORKDIR /usr/src/app
COPY ["views/","views/"]
COPY ["statics/","statics/"]
COPY ["app", "/usr/src/app"]
EXPOSE 2666
ENTRYPOINT ["./app"]
