# ETAPA 1: Compilación (Imagen pesada con herramientas de Go)
FROM golang:1.23-alpine AS builder
WORKDIR /app
# Copiamos archivos de módulos primero para aprovechar el caché de Docker
COPY go.mod go.sum ./
RUN go mod download
# Copiamos el resto del código y archivos estáticos
COPY . .
# Compilamos el binario para Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main-app .

# ETAPA 2: Ejecución (Imagen ligera)
FROM alpine:latest
WORKDIR /root/
# Copiamos el binario y los archivos web desde la etapa anterior
COPY --from=builder /app/main-app .
COPY --from=builder /app/index.html .
COPY --from=builder /app/styles.css .

# Exponemos el puerto (ajusta si usas uno diferente al 8080)
EXPOSE 8080

CMD ["./main-app"]