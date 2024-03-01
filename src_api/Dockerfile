# Use a imagem oficial do Go como base
FROM golang:1.20-alpine AS builder

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o código fonte para o diretório de trabalho
COPY . .

# Compile o aplicativo
RUN go build -o api_rinha_de_backend_2024_q1 .

# Use uma imagem leve como base para reduzir o tamanho final da imagem
FROM alpine:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /root/

# Copie o executável compilado do estágio anterior
COPY --from=builder /app/api_rinha_de_backend_2024_q1 .

# Defina a variável de ambiente PORT
ENV PORT=8080

# Exponha a porta necessária para o aplicativo
EXPOSE $PORT

# Comando para executar o aplicativo quando o contêiner for iniciado
CMD ["./api_rinha_de_backend_2024_q1"]
