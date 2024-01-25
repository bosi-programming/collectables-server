# Use a imagem oficial do Go como base
FROM golang:1.21.6-alpine3.18 as build

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /go-server

# Download Go modules
COPY . .
RUN go mod download

# Construa o aplicativo Go sem debug
RUN CGO_ENABLED=0 go build -ldflags '-w -s' -o go-web-server

From alpine:3.18
WORKDIR /go-server
COPY --from=build /go-server .
EXPOSE 3000
CMD ["./go-web-server"]
