FROM alpine:3.7
ADD static /bin/static
ENTRYPOINT ["/bin/static"]

