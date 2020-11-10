FROM alpine

#ENV CONFIG_SERVER http://localhost:8888
#ENV PROFILE test
#ENV BRANCH blog

COPY menuservice-linux-amd64 /
COPY healthchecker-linux-amd64 /
COPY config.yml /

HEALTHCHECK --interval=10s --timeout=10s --start-period=10s --retries=3 CMD [ "/healthchecker-linux-amd64","http://localhost:2604/menu" ] || exit 1

EXPOSE 2604 

ENTRYPOINT ["/menuservice-linux-amd64"] 
