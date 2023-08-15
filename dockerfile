FROM golang:1.20-alpine as dependencies

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM dependencies AS build
COPY . ./
RUN CGO_ENABLED=0 go build -o /main -ldflags="-w -s"

FROM go:1.20-apline
COPY --from=build /main /main
CMD ["/main"]
