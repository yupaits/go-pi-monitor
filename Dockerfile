FROM armhf/alpine:3.5

LABEL maintainer="yupaits"

ADD main app/
ADD ui/dist app/ui/dist
ADD config.toml app/

EXPOSE 80
CMD app/main