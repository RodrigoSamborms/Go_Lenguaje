FROM golang:1.23-alpine

WORKDIR /app

# Al usar el contexto de la raíz, Docker encuentra go.mod aquí
COPY go.mod go.sum ./
RUN go mod download

# Copiamos solo el contenido de la carpeta correspondiente
COPY frontend/ . 
# (Y en el otro Dockerfile: COPY frontend/ .)

RUN go build -o main_web .

EXPOSE 8080

CMD ["./main_web"]