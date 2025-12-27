CREATE DATABASE IF NOT EXISTS proyecto_go;
CREATE USER IF NOT EXISTS 'usuariodb'@'%' IDENTIFIED BY '1234';
GRANT ALL PRIVILEGES ON proyecto_go.* TO 'usuariodb'@'%';
FLUSH PRIVILEGES;
USE proyecto_go;

CREATE TABLE IF NOT EXISTS usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    activo BOOLEAN DEFAULT TRUE
);

-- Insertar algunos datos de prueba
INSERT INTO usuarios (nombre) VALUES ('Rodrigo'), ('Sambo'), ('Go Developer');