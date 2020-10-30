FROM alpine

#ENV CONFIG_SERVER http://localhost:8888
#ENV PROFILE test
#ENV BRANCH blog

COPY postsservice-linux-amd64 /
COPY healthchecker-linux-amd64 /
COPY config.yml /

HEALTHCHECK --interval=10s --timeout=10s --start-period=10s --retries=3 CMD [ "/healthchecker-linux-amd64","http://localhost:2600/posts" ] || exit 1

EXPOSE 2600 

ENTRYPOINT ["/postsservice-linux-amd64"] 
#ENTRYPOINT  /musicservice-linux-amd64 -configServer ${CONFIG_SERVER} -profile ${PROFILE} -branch ${BRANCH} 
