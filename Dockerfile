FROM alpine:3.1
MAINTAINER Alberto Vila <alberto.vila@protonmail.com>
ADD rest-api /usr/bin/rest-api
ENTRYPOINT ["rest-api"]
