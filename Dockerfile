FROM golang
WORKDIR /app
COPY server.go /app/
COPY web /app/web/
COPY data /app/data/
RUN go build server.go
RUN pwd
RUN ls -la /app
EXPOSE 8080
CMD ["./server"]