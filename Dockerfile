FROM scratch
COPY respond /usr/bin/respond
ENTRYPOINT ["/usr/bin/respond"]
EXPOSE 8080
CMD ["-port", "8080"]
