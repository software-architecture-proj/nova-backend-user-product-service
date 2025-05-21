# syntax=docker/dockerfile:1.4 
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o user_product_service ./main.go

FROM alpine:latest
RUN apk add --no-cache postgresql-client
WORKDIR /root/
COPY --from=builder /app/user_product_service .
CMD ["./user_product_service"]


# To build the Docker image, run:

#echo "github_pat_xxxxxxxxxxxxxxxxxxx" > ~/.github-token  
#chmod 600 ~/.github-token   
#DOCKER_BUILDKIT=1 docker build \
#  --secret id=github_token,src=$HOME/.github-token \
#  -t user_product_service:x .