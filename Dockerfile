# Build the Go API
FROM golang:latest AS builder
ADD . /app
WORKDIR /app/application-server
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
ADD . /app
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN chmod +x ./main
EXPOSE 8080
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools
RUN pip3 install -r vulnerability/query_github/requirements.txt
CMD ./main & python3 vulnerability/query_github/fetch_github_security_vulnerabilities.py

