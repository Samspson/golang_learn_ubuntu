FROM plugins/base:multiarch

LABEL maintainer="wxp <wxp_sampson@163.com>" \
    org.label-schema.name="Drone Workshop" \
    org.label-schema.vendor="wxp" \
    org.label-schema.schema-version="1.0"

ADD release/linux/amd64/helloworld /bin/

ENTRYPOINT [ "/bin/helloworld" ]