# Build the Go API
FROM golang:latest AS builder
ADD . /app
WORKDIR /app/application-server
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# Build python
FROM python:alpine AS python_builder
RUN pip install -r /app/vulnerability/query_github/requirements.txt
CMD python fetch_github_security_vulnerabilities.py &

# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN chmod +x ./main
EXPOSE 8080
CMD ./main
