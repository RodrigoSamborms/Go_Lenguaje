<?php
// borrar.php - Eliminar un usuario (marcar como inactivo)

include 'config.php';

$id = isset($_GET['id']) ? (int)$_GET['id'] : 0;

if ($id <= 0) {
    header("Location: listar.php");
    exit;
}

// Marcar como inactivo en lugar de eliminar (soft delete)
$stmt = $conn->prepare("UPDATE usuarios SET activo = 0 WHERE id = ? AND activo = 1");
$stmt->bind_param("i", $id);

if ($stmt->execute()) {
    if ($stmt->affected_rows > 0) {
        header("Location: listar.php?success=Usuario eliminado correctamente");
    } else {
        header("Location: listar.php?error=Usuario no encontrado");
    }
} else {
    header("Location: listar.php?error=Error al eliminar");
}

$stmt->close();
$conn->close();
?>
