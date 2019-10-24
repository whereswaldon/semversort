FROM scratch
COPY semversort /
ENTRYPOINT ["/semversort"]
