# Arquitectura Monolítica - PHP + MariaDB

## 📋 Descripción

Esta carpeta contiene una implementación **monolítica tradicional** de una aplicación CRUD para gestión de nombres usando:
- **PHP 8.2** - Lenguaje de servidor
- **Apache** - Servidor web
- **MariaDB 10.11** - Base de datos relacional
- **Docker** - Containerización

Esta arquitectura contrasta con las versiones en microservicios de las carpetas `Microservicio/` y `V2_Microservicio/`.

---

## 🏗️ Estructura del Proyecto

```
Monolotico/
├── index.php              # Página principal
├── listar.php             # Listar usuarios
├── agregar.php            # Crear usuarios
├── editar.php             # Modificar usuarios
├── borrar.php             # Eliminar usuarios
├── config.php             # Configuración de BD
├── styles.css             # Estilos CSS
├── init.sql               # Script de inicialización de BD
├── docker-compose.yml     # Orquestación con Docker
└── README.md              # Este archivo
```

---

## 🚀 Cómo Ejecutar

### Opción 1: Con Docker Compose (Recomendado)

```bash
cd Monolotico/

# Iniciar los servicios
docker-compose up -d

# Esperar 10 segundos a que se inicialice la BD
# Luego acceder a: http://localhost:8080
```

**Detener:**
```bash
docker-compose down
```

**Ver logs:**
```bash
docker-compose logs -f web    # Logs del PHP
docker-compose logs -f db     # Logs de MariaDB
```

---

### Opción 2: Ejecución Local

**Requisitos:**
- PHP 8.2+ con extensión mysqli
- Apache o PHP built-in server
- MariaDB/MySQL corriendo localmente

**Pasos:**

```bash
# 1. Inicializar base de datos
mysql -u root -p < init.sql

# 2. Actualizar config.php si es necesario
# cambiar "localhost" a la ruta de tu servidor

# 3. Opción A: Con PHP built-in server
php -S localhost:8080

# 3. Opción B: Con Apache
# Copiar archivos a /var/www/html y acceder a http://localhost
```

---

## 🎯 Funcionalidades CRUD

### ✅ Listar Usuarios
- **Archivo:** `listar.php`
- **URL:** `http://localhost:8080/listar.php`
- **Descripción:** Muestra todos los usuarios activos con opciones de editar y borrar

### ✅ Crear Usuario
- **Archivo:** `agregar.php`
- **URL:** Formulario en `index.php` o acceso directo a `agregar.php`
- **Descripción:** Añade un nuevo usuario a la base de datos

### ✅ Editar Usuario
- **Archivo:** `editar.php`
- **Descripción:** Modifica el nombre de un usuario existente

### ✅ Eliminar Usuario
- **Archivo:** `borrar.php`
- **Descripción:** Marca un usuario como inactivo (soft delete)

---

## 📊 Comparación: Monolítica vs Microservicios

| Aspecto | Monolítica (PHP) | Microservicios (Go) |
|---------|------------------|-------------------|
| **Ubicación** | Monolotico/ | Microservicio/, V2_Microservicio/ |
| **Lenguaje** | PHP | Go |
| **Base de Datos** | Integrada | Separada en servicio |
| **Frontend** | Mismo archivo PHP | Servicio separado |
| **API** | No (acoplada) | REST JSON desacoplada |
| **Escalabilidad** | Vertical | Horizontal |
| **Despliegue** | Todo junto | Servicios independientes |
| **Orquestación** | Docker Compose | Kubernetes |

---

## 🔧 Configuración de Base de Datos

### Credenciales Predeterminadas

```
Host: localhost (o "db" con Docker)
Usuario: usuariodb
Contraseña: 1234
Base de datos: proyecto_go
Puerto: 3306
```

### Modificar Credenciales

1. Editar `docker-compose.yml`:
   ```yaml
   MYSQL_USER: otro_usuario
   MYSQL_PASSWORD: otra_contraseña
   ```

2. Editar `config.php`:
   ```php
   $username = "otro_usuario";
   $password = "otra_contraseña";
   ```

3. Editar `init.sql`:
   ```sql
   CREATE USER 'otro_usuario'@'%' IDENTIFIED BY 'otra_contraseña';
   ```

---

## 📁 Estructura de Base de Datos

### Tabla: usuarios

```sql
CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    activo BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Campos:**
- `id` - Identificador único
- `nombre` - Nombre del usuario (máx 100 caracteres)
- `activo` - Estado (1=activo, 0=inactivo) para soft deletes
- `created_at` - Timestamp de creación

---

## 🛡️ Características de Seguridad

✅ **Prepared Statements** - Previenen SQL injection  
✅ **htmlspecialchars()** - Escapan caracteres especiales  
✅ **Validación de entrada** - Verifican datos antes de procesar  
✅ **Soft Deletes** - No eliminar, solo marcar como inactivo  
✅ **Set Charset UTF-8** - Soporte completo de caracteres especiales  

---

## 🐛 Troubleshooting

### Error: "Error de conexión"
```
Solución: Asegurate que MariaDB esté corriendo
- Con Docker: docker-compose up -d db
- Localmente: sudo systemctl start mariadb
```

### Error: "Base de datos no encontrada"
```
Solución: Ejecutar init.sql
- Con Docker: Automático en el primer inicio
- Localmente: mysql -u root -p < init.sql
```

### Error: "mysqli extension not loaded"
```
Solución: Instalar extensión MySQLi
- Docker: Ya incluido en docker-compose.yml
- Local: docker-php-ext-install mysqli
```

---

## 📚 Recursos

- [PHP Documentation](https://www.php.net/docs.php)
- [MariaDB Documentation](https://mariadb.com/kb/en/documentation/)
- [Docker Documentation](https://docs.docker.com/)

---

## 🎓 Propósito Educativo

Este proyecto es un ejemplo educativo que demuestra:
- Desarrollo web tradicional monolítico
- Arquitectura acoplada Frontend + Backend + BD
- Limitaciones de escalabilidad vertical
- Cómo esta arquitectura evolucionó a microservicios (ver carpetas Microservicio/ y V2_Microservicio/)

---

## 🔜 Siguientes Pasos

Para ver cómo evolucionó esta arquitectura:
1. Revisar `Microservicio/` - Versión en Go con Docker Compose
2. Revisar `V2_Microservicio/` - Versión con Kubernetes
3. Ver `README.md` en raíz para entender la evolución completa

---

**Estado:** ✅ Funcional  
**Última actualización:** 27 de Diciembre, 2025
