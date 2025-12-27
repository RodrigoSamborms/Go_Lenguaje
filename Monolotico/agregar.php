<?php
// agregar.php - Crear un nuevo usuario

include 'config.php';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $nombre = trim($_POST['nombre'] ?? '');
    
    if (empty($nombre)) {
        $error = "El nombre no puede estar vacío";
    } else if (strlen($nombre) > 100) {
        $error = "El nombre no puede exceder 100 caracteres";
    } else {
        // Preparar consulta para evitar SQL injection
        $stmt = $conn->prepare("INSERT INTO usuarios (nombre, activo) VALUES (?, 1)");
        $stmt->bind_param("s", $nombre);
        
        if ($stmt->execute()) {
            header("Location: listar.php?success=Usuario creado correctamente");
            exit;
        } else {
            $error = "Error al crear el usuario: " . $stmt->error;
        }
        $stmt->close();
    }
}

$conn->close();
?>
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Agregar Nombre</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <div class="seccion">
        <h1>Agregar Nuevo Nombre</h1>
        
        <?php if (isset($error)): ?>
            <div class="error"><?php echo htmlspecialchars($error); ?></div>
        <?php endif; ?>
        
        <form action="agregar.php" method="POST">
            <input type="text" name="nombre" placeholder="Escribe un nombre" required autofocus>
            <button type="submit">Guardar</button>
        </form>
        
        <hr>
        <a href='index.php' class='btn' style='background-color: #6c757d;'>← Regresar</a>
    </div>
</body>
</html>
