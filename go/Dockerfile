FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .
RUN go build -o loanbuddy ./cmd/loanbuddy

FROM gcr.io/distroless/static
COPY --from=build /app/loanbuddy /
ENTRYPOINT ["/loanbuddy"]