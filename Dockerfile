FROM golang:alpine AS build

WORKDIR /app
COPY . .
RUN go build -o uno_server

FROM alpine

WORKDIR /app
COPY --from=build /app/uno_server .
COPY client client

CMD ./uno_server