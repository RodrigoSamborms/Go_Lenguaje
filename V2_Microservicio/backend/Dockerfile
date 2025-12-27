FROM golang:1.23-alpine

WORKDIR /app

# Al usar el contexto de la raíz, Docker encuentra go.mod aquí
COPY go.mod go.sum ./
RUN go mod download

# Copiamos solo el contenido de la carpeta correspondiente
COPY backend/ . 
# (Y en el otro Dockerfile: COPY frontend/ .)

RUN go build -o main_api .

EXPOSE 8081

CMD ["./main_api"]