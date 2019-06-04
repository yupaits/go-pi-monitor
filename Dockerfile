FROM armhf/alpine:3.5

LABEL maintainer="yupaits"

RUN ls /app/
RUN ls /app/ui/dist

ADD main /app/
ADD ui/dist /app/ui/dist
ADD config.toml /app/

EXPOSE 80
ENTRYPOINT ["/app/main"]