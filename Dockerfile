FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/topix /opt/bin/topix
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/topix"]

