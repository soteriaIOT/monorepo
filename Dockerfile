# Build the Go API
FROM golang:latest AS builder
ADD . /app
WORKDIR /app/application-server
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# Build python
FROM python:alpine AS python_builder
COPY --from=builder /app/vulnerability/query_github ./
RUN pip3 install -r requirements.txt
CMD python3 fetch_github_security_vulnerabilities.py &

# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN chmod +x ./main
EXPOSE 8080
CMD ./main
