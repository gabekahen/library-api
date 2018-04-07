FROM alpine:latest
LABEL maintainer "gabe@gabekahen.com"

EXPOSE 8080/tcp


ENV LIBRARYAPI_DB_USER "root"
ENV LIBRARYAPI_DB_PASS ""
ENV LIBRARYAPI_DB_PROT "tcp"
ENV LIBRARYAPI_DB_HOST "mysql"

ADD library_api /

RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

HEALTHCHECK --timeout=5s CMD curl -f http://localhost:8080/debug/health || exit 1

CMD ["/library_api"]
