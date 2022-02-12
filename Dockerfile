FROM alpine
COPY login-service /
ENTRYPOINT ["/login-service"]
EXPOSE 6789