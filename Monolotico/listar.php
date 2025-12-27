<?php
// listar.php - Mostrar todos los usuarios

include 'config.php';

// Obtener usuarios activos
$sql = "SELECT id, nombre FROM usuarios WHERE activo = 1 ORDER BY id DESC";
$result = $conn->query($sql);

?>
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Listado de Nombres</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <div class="seccion">
        <h1>Usuarios Activos (Monolítica - PHP + MariaDB)</h1>
        
        <?php
        if ($result && $result->num_rows > 0) {
            echo "<ul>";
            while ($row = $result->fetch_assoc()) {
                $id = $row['id'];
                $nombre = htmlspecialchars($row['nombre']);
                echo "<li>";
                echo "<span class='nombre-usuario'>[$id] $nombre</span>";
                echo "<div class='acciones'>";
                echo "<a href='borrar.php?id=$id' onclick=\"return confirm('¿Eliminar a $nombre?')\" class='btn btn-borrar'>Borrar</a>";
                echo "<form action='editar.php' method='POST' style='display:inline;'>";
                echo "<input type='hidden' name='id' value='$id'>";
                echo "<input type='text' name='nuevoNombre' placeholder='Nuevo nombre' required>";
                echo "<button type='submit' class='btn btn-editar'>Editar</button>";
                echo "</form>";
                echo "</div>";
                echo "</li>";
            }
            echo "</ul>";
            echo "<p><strong>Total de usuarios: " . $result->num_rows . "</strong></p>";
        } else {
            echo "<p style='color: #666;'>No hay usuarios registrados.</p>";
        }
        ?>
        
        <hr>
        <a href='index.php' class='btn' style='background-color: #6c757d;'>← Regresar</a>
    </div>
</body>
</html>

<?php
$conn->close();
?>
