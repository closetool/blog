FROM alpine

COPY fileservice-linux-amd64 /
COPY healthchecker-linux-amd64 /
COPY config.yml /

HEALTHCHECK --interval=10s --timeout=10s --start-period=10s --retries=3 CMD [ "/healthchecker-linux-amd64","http://localhost:2607/file" ] || exit 1

EXPOSE 2607 

ENTRYPOINT ["/fileservice-linux-amd64"] 