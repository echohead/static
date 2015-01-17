FROM echohead/nobody
ADD static /bin/static
USER nobody
ENTRYPOINT [ "/bin/static" ]
