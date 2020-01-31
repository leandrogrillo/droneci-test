FROM alpine
ADD ./droneci-test /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/droneci-test