<?php
// config.php - Configuración de conexión a la base de datos

$servername = "db";  // Nombre del servicio de base de datos en Docker
$username = "usuariodb";
$password = "1234";
$dbname = "proyecto_go";

// Crear conexión
$conn = new mysqli($servername, $username, $password, $dbname);

// Verificar conexión
if ($conn->connect_error) {
    die("Error de conexión: " . $conn->connect_error);
}

// Establecer charset
$conn->set_charset("utf8mb4");

?>
