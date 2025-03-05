# syntax=docker/dockerfile:1

FROM node:lts-alpine
WORKDIR /app
COPY . .
CMD ["go run main.go"]
EXPOSE 30001