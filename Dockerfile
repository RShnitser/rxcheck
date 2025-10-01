FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY server /bin/rxcheck/server

CMD ["/bin/rxcheck/server"]