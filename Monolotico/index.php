<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Monolítica - Gestión de Nombres</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <h1>Gestión de Nombres (PHP + MariaDB Monolítica)</h1>
    <div class="seccion">
        <h3>Agregar Nombre</h3>
        <form action="agregar.php" method="POST">
            <input type="text" name="nombre" placeholder="Escribe un nombre" required>
            <button type="submit">Guardar</button>
        </form>
    </div>
    <div class="seccion">
        <h3>Ver Nombres Guardados</h3>
        <a href="listar.php">
            <button style="background-color: #28a745;">Ver Listado de Nombres</button>
        </a>
    </div>
</body>
</html>
