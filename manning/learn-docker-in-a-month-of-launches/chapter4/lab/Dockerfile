FROM diamol/golang as builder

COPY main.go .
RUN go build -o /server
RUN chmod +x /server

# app
FROM diamol/base

EXPOSE 80
ENV USER=yakushou730
CMD ["/web/server"]

WORKDIR web
COPY --from=builder /server .

COPY index.html .
