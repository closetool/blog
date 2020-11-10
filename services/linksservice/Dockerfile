FROM alpine

COPY linksservice-linux-amd64 /
COPY healthchecker-linux-amd64 /
COPY config.yml /

HEALTHCHECK --interval=10s --timeout=10s --start-period=10s --retries=3 CMD [ "/healthchecker-linux-amd64","http://localhost:2606/link" ] || exit 1

EXPOSE 2606 

ENTRYPOINT ["/linksservice-linux-amd64"] 