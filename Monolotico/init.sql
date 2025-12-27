CREATE DATABASE IF NOT EXISTS proyecto_go;

-- Crear usuario con permisos
DROP USER IF EXISTS 'usuariodb'@'%';
CREATE USER IF NOT EXISTS 'usuariodb'@'%' IDENTIFIED BY '1234';
GRANT ALL PRIVILEGES ON proyecto_go.* TO 'usuariodb'@'%';
FLUSH PRIVILEGES;

USE proyecto_go;

-- Crear tabla de usuarios
CREATE TABLE IF NOT EXISTS usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    activo BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insertar datos de prueba
INSERT INTO usuarios (nombre, activo) VALUES ('Rodrigo', 1);
INSERT INTO usuarios (nombre, activo) VALUES ('Sambo', 1);
INSERT INTO usuarios (nombre, activo) VALUES ('Go Developer', 1);
