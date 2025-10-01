FROM debian:stable-slim

COPY server /bin/rxcheck/server
COPY static/auth.js /bin/rxcheck/static/auth.js
COPY static/htmx.min.js /bin/rxcheck/static/htmx.min.js
COPY static/styles.css /bin/rxcheck/static/styles.css

CMD ["/bin/rxcheck/server"]