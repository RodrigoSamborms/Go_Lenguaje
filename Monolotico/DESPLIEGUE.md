# 🚀 Guía de Despliegue - Dos Enfoques

**Fecha:** 27 de Diciembre, 2025  
**Propósito:** Ejecutar la aplicación monolítica en dos entornos diferentes  
**Entorno Principal:** WSL Debian en Windows con Docker

---

## 📍 Dos Opciones de Ejecución

### Opción A: Raspberry Pi + Apache Web Server
Para producción en hardware embebido (sin Docker)

### Opción B: Máquina x86 Local + Docker Compose
Para desarrollo y pruebas locales (ejecutándose en WSL Debian)

---

## 🍓 OPCIÓN A: Raspberry Pi + Apache

### 📋 Requisitos

En la **Raspberry Pi:**
- ✅ Apache Web Server corriendo
- ✅ PHP 8.2+ instalado
- ✅ MariaDB/MySQL corriendo
- ✅ Acceso SSH activado

### 📝 Pasos de Despliegue

#### 1. Preparar Archivos Locales (En tu PC Windows)

```bash
cd Go_Lenguaje/Monolotico/

# Archivos necesarios a copiar:
# - agregar.php
# - borrar.php
# - editar.php
# - index.php
# - listar.php
# - config.php
# - styles.css
```

#### 2. Copiar Archivos a la Raspberry Pi

**En PowerShell o Terminal:**

```bash
# Reemplaza pi@192.168.1.X con la IP de tu RPi
$RPI_IP = "192.168.1.100"  # Cambia a tu IP
$RPI_USER = "pi"           # O tu usuario en RPi

# Copiar archivos PHP
scp agregar.php ${RPI_USER}@${RPI_IP}:/var/www/html/
scp borrar.php ${RPI_USER}@${RPI_IP}:/var/www/html/
scp editar.php ${RPI_USER}@${RPI_IP}:/var/www/html/
scp index.php ${RPI_USER}@${RPI_IP}:/var/www/html/
scp listar.php ${RPI_USER}@${RPI_IP}:/var/www/html/
scp config.php ${RPI_USER}@${RPI_IP}:/var/www/html/
scp styles.css ${RPI_USER}@${RPI_IP}:/var/www/html/
```

#### 3. Configurar Base de Datos en RPi

**En SSH en la Raspberry Pi:**

```bash
# Conectar a MySQL
mysql -u root -p

# Ejecutar el script init.sql
mysql -u root -p < init.sql

# O copiar y ejecutar manualmente:
# CREATE DATABASE proyecto_go;
# CREATE USER 'usuariodb'@'%' IDENTIFIED BY '1234';
# GRANT ALL PRIVILEGES ON proyecto_go.* TO 'usuariodb'@'%';
# FLUSH PRIVILEGES;
```

#### 4. Configurar config.php en RPi

**En la RPi, editar `/var/www/html/config.php`:**

```php
<?php
// Detectar si estamos en RPi
$is_rpi = (gethostname() === 'raspberrypi' || 
           file_exists('/etc/os-release') && strpos(file_get_contents('/etc/os-release'), 'Debian') !== false);

if ($is_rpi) {
    // Configuración para Raspberry Pi
    $servername = "localhost";
} else {
    // Fallback para otros entornos
    $servername = "localhost";
}

$username = "usuariodb";
$password = "1234";
$dbname = "proyecto_go";

$conn = new mysqli($servername, $username, $password, $dbname);

if ($conn->connect_error) {
    die("Error de conexión: " . $conn->connect_error);
}

$conn->set_charset("utf8mb4");
?>
```

#### 5. Verificar Permisos en RPi

```bash
# En SSH en la RPi:
sudo chown www-data:www-data /var/www/html/*.php
sudo chown www-data:www-data /var/www/html/styles.css
sudo chmod 755 /var/www/html/*.php
sudo chmod 644 /var/www/html/styles.css
```

#### 6. Reiniciar Apache en RPi

```bash
sudo systemctl restart apache2
```

#### 7. Acceder a la Aplicación

```
Desde tu navegador:
http://raspberrypi.local
o
http://192.168.1.100  (reemplaza con tu IP)
```

---

### ⚙️ Troubleshooting RPi

**Problema: "Error de conexión a base de datos"**
```bash
# Verificar que MariaDB está corriendo
sudo systemctl status mariadb

# Iniciar si no está activo
sudo systemctl start mariadb
```

**Problema: "Permiso denegado"**
```bash
# Cambiar permisos
sudo chown -R www-data:www-data /var/www/html/
sudo chmod -R 755 /var/www/html/
```

**Problema: "Página en blanco"**
```bash
# Ver logs de Apache
tail -f /var/log/apache2/error.log
```

---

## 💻 OPCIÓN B: Máquina x86 Local + Docker Compose

### 📋 Requisitos

En tu **máquina Windows/Linux/macOS:**
- ✅ Docker Desktop instalado
- ✅ Docker Compose instalado
- ✅ Puertos 8080 y 3306 disponibles

### 📝 Pasos de Despliegue

#### 1. Verificar Requisitos

```bash
docker --version
docker-compose --version
```

#### 2. Navegar a la Carpeta

```bash
cd Go_Lenguaje/Monolotico
```

#### 3. Levantar los Servicios

```bash
# Iniciar servicios
docker-compose up -d

# Esperar 10-15 segundos a que MariaDB se inicialice
sleep 15

# Verificar estado
docker-compose ps
```

#### 4. Acceder a la Aplicación

```
http://localhost:8080
```

---

### ⚙️ Troubleshooting Docker

**Problema: "Puerto 8080 ya está en uso"**
```bash
# Cambiar puerto en docker-compose.yml:
# ports:
#   - "8081:80"  (en lugar de 8080:80)

# Luego acceder a: http://localhost:8081
```

**Problema: "Base de datos no se inicializa"**
```bash
# Ver logs
docker-compose logs db

# Eliminar volumen y reintentar
docker-compose down -v
docker-compose up -d
```

**Problema: "Conexión rechazada"**
```bash
# Esperar más tiempo y verificar
sleep 20
docker-compose ps

# Si mariadb sigue sin estar Ready:
docker-compose logs db
```

---

## 📊 Comparación Rápida

| Aspecto | Raspberry Pi | Docker Local |
|---------|-------------|--------------|
| **Requisitos** | Apache, PHP, MySQL | Docker Desktop |
| **Instalación** | Manual | Automática |
| **Performance** | Optimizado para RPi | Completo en x86 |
| **Persistencia** | Base local RPi | Volumen Docker |
| **Desarrollo** | Producción | Testing/Desarrollo |
| **Facilidad** | Media | Alta |
| **Tiempo setup** | 15-20 min | 2-3 min |

---

## 🔄 Flujo de Trabajo Recomendado

```
1. DESARROLLO LOCAL (Docker x86)
   ├─ docker-compose up -d
   ├─ Probar cambios
   └─ Iterar rápidamente
   
2. VERSIONAMIENTO (Git)
   ├─ git add Monolotico/
   └─ git commit
   
3. DESPLIEGUE EN RPi
   ├─ scp archivos *.php
   ├─ scp styles.css
   └─ mysql < init.sql (una sola vez)
   
4. MONITOREO
   ├─ Acceder a http://raspberrypi.local
   └─ Ver logs: tail -f /var/log/apache2/error.log
```

---

## 📝 Arquivos Necesarios por Entorno

### Para Raspberry Pi (Copiar estos 7)
```
✅ agregar.php
✅ borrar.php
✅ editar.php
✅ index.php
✅ listar.php
✅ config.php
✅ styles.css
```

### Para Docker Local (Todo, pero específicamente)
```
✅ Todos los archivos PHP
✅ styles.css
✅ docker-compose.yml ← REQUERIDO
✅ init.sql ← Para inicializar BD
✅ php.ini ← Para configuración
```

---

## 🚀 Cómo Actualizar la Aplicación

### En Raspberry Pi
```bash
# Simplemente copiar nuevos archivos
scp archivo_actualizado.php pi@192.168.1.100:/var/www/html/

# Apache lo recoge automáticamente
```

### En Docker Local
```bash
# Editar archivo localmente
# Docker refleja cambios inmediatamente

# Si cambias config de BD:
docker-compose restart db
```

---

## 🎯 Casos de Uso

### Caso 1: Quiero Probar Localmente Primero
```
1. docker-compose up -d (en x86)
2. Desarrollo en localhost:8080
3. Luego copiar a RPi cuando esté estable
```

### Caso 2: Quiero Ir Directo a Producción RPi
```
1. Preparar archivos en Windows
2. scp a RPi
3. Ejecutar init.sql en RPi
4. Acceder directamente
```

### Caso 3: Quiero Ambos Ambientes Sincronizados
```
1. Git como "fuente de verdad"
2. git pull en RPi regularmente
3. git push con cambios locales
4. Ambos usan mismo código
```

---

## ✅ Checklist de Despliegue

### Para Raspberry Pi
- [ ] SSH activado en RPi
- [ ] Apache corriendo
- [ ] PHP 8.2+ instalado
- [ ] MariaDB corriendo
- [ ] Archivos copiados a /var/www/html
- [ ] Base de datos creada (init.sql ejecutado)
- [ ] Permisos configurados
- [ ] Apache reiniciado
- [ ] Acceso verificado

### Para Docker Local
- [ ] Docker Desktop activo
- [ ] docker-compose up -d ejecutado
- [ ] Esperar 15 segundos
- [ ] docker-compose ps muestra todos Running
- [ ] http://localhost:8080 accesible
- [ ] CRUD completo funciona

---

## 📖 Documentación Relacionada

- `README.md` - Descripción general
- `TESTING.md` - Casos de prueba
- `docker-compose.yml` - Configuración Docker
- `init.sql` - Script de BD

---

**Creado:** 27 de Diciembre, 2025  
**Propósito:** Guía de despliegue en dos arquitecturas  
**Estado:** ✅ Listo para usar
