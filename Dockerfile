FROM alpine
ADD static /bin/static
ENTRYPOINT ["/bin/static"]

