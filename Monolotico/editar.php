<?php
// editar.php - Editar un usuario existente

include 'config.php';

$id = isset($_GET['id']) ? (int)$_GET['id'] : (isset($_POST['id']) ? (int)$_POST['id'] : 0);

if ($id <= 0) {
    header("Location: listar.php");
    exit;
}

// Si es POST, actualizar el usuario
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $nuevoNombre = trim($_POST['nuevoNombre'] ?? '');
    
    if (empty($nuevoNombre)) {
        $error = "El nombre no puede estar vacío";
    } else if (strlen($nuevoNombre) > 100) {
        $error = "El nombre no puede exceder 100 caracteres";
    } else {
        $stmt = $conn->prepare("UPDATE usuarios SET nombre = ? WHERE id = ? AND activo = 1");
        $stmt->bind_param("si", $nuevoNombre, $id);
        
        if ($stmt->execute()) {
            if ($stmt->affected_rows > 0) {
                header("Location: listar.php?success=Usuario actualizado correctamente");
                exit;
            } else {
                $error = "Usuario no encontrado o no activo";
            }
        } else {
            $error = "Error al actualizar: " . $stmt->error;
        }
        $stmt->close();
    }
}

// Obtener datos del usuario para mostrar
$stmt = $conn->prepare("SELECT id, nombre FROM usuarios WHERE id = ? AND activo = 1");
$stmt->bind_param("i", $id);
$stmt->execute();
$result = $stmt->get_result();

if ($result->num_rows === 0) {
    header("Location: listar.php");
    exit;
}

$usuario = $result->fetch_assoc();
$stmt->close();
$conn->close();
?>
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Editar Nombre</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <div class="seccion">
        <h1>Editar Nombre</h1>
        <p>Usuario ID: <strong><?php echo $usuario['id']; ?></strong></p>
        <p>Nombre actual: <strong><?php echo htmlspecialchars($usuario['nombre']); ?></strong></p>
        
        <?php if (isset($error)): ?>
            <div class="error"><?php echo htmlspecialchars($error); ?></div>
        <?php endif; ?>
        
        <form action="editar.php" method="POST">
            <input type="hidden" name="id" value="<?php echo $usuario['id']; ?>">
            <input type="text" name="nuevoNombre" placeholder="Nuevo nombre" value="<?php echo htmlspecialchars($usuario['nombre']); ?>" required autofocus>
            <button type="submit">Actualizar</button>
        </form>
        
        <hr>
        <a href='listar.php' class='btn' style='background-color: #6c757d;'>← Regresar</a>
    </div>
</body>
</html>
