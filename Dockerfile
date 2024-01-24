# Use a imagem oficial do Go como base
FROM golang:1.21-alpine

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /go-server

# Download Go modules
COPY . .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Construa o aplicativo Go
RUN go build -o go-server

# Exponha a porta em que o aplicativo será executado
EXPOSE 3000

# Comando padrão para executar o aplicativo quando o contêiner for iniciado
CMD ["./go-server"]
