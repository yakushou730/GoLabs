FROM diamol/node

CMD ["node", "/web-ping/app.js"]

ENV TARGET="docker.com" \
    METHOD="HEAD" \
    INTERVAL="1000"

WORKDIR /web-ping
COPY app.js .

