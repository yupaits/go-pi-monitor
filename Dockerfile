FROM armhf/alpine:3.5

LABEL maintainer="yupaits"

ADD main app/go-pi-monitor
ADD ui/dist app/ui/dist
ADD config.toml app/config.toml

EXPOSE 80
CMD app/go-pi-monitor