FROM armhf/alpine:3.5

LABEL maintainer="yupaits"

ADD main /app/
ADD ui/dist /app/ui/dist
ADD config.toml /app/

RUN chmod +x /app/main

EXPOSE 80
ENTRYPOINT ["/app/main"]