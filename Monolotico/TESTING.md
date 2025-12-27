# 🧪 Guía de Prueba - Arquitectura Monolítica PHP

## ⚡ Inicio Rápido

### 1. Levantar la Aplicación

```bash
# Desde la carpeta Monolotico/
docker-compose up -d

# Esperar 10-15 segundos mientras se inicializa MariaDB
```

### 2. Verificar que Todo Está Corriendo

```bash
# Verificar contenedores
docker-compose ps

# Resultado esperado:
# NAME              STATUS      PORTS
# proyecto_go_web   Up 10 min   0.0.0.0:8080->80/tcp
# proyecto_go_db    Up 10 min   0.0.0.0:3306->3306/tcp
```

### 3. Acceder a la Aplicación

Abrir navegador en: **http://localhost:8080**

---

## 📝 Casos de Prueba

### Test 1: Verificar Página Principal

**Acción:** Acceder a `http://localhost:8080`

**Resultado Esperado:**
- ✅ Se muestra formulario "Agregar Nombre"
- ✅ Se muestra botón "Ver Listado de Nombres"
- ✅ Estilos CSS se aplican correctamente

---

### Test 2: Listar Usuarios Existentes

**Acción:** Hacer click en "Ver Listado de Nombres"

**Resultado Esperado:**
- ✅ Se muestran 3 usuarios de prueba:
  - [1] Rodrigo
  - [2] Sambo
  - [3] Go Developer
- ✅ Cada usuario tiene botones "Borrar" y "Editar"
- ✅ Se muestra total: "Total de usuarios: 3"

---

### Test 3: Crear Nuevo Usuario

**Acción:**
1. Desde la página principal
2. Ingresar nombre: "Test Usuario"
3. Click en "Guardar"

**Resultado Esperado:**
- ✅ Redirige a listado
- ✅ Nuevo usuario aparece al inicio de la lista
- ✅ Total de usuarios aumenta a 4

---

### Test 4: Editar Usuario

**Acción:**
1. Desde listado
2. Hacer click en "Editar" para "Rodrigo"
3. Cambiar nombre a "Rodrigo Modificado"
4. Click en "Actualizar"

**Resultado Esperado:**
- ✅ Redirige a listado
- ✅ Nombre actualizado: [1] Rodrigo Modificado
- ✅ Los demás usuarios no cambian

---

### Test 5: Eliminar Usuario

**Acción:**
1. Desde listado
2. Hacer click en "Borrar" para "Sambo"
3. Confirmar en diálogo del navegador

**Resultado Esperado:**
- ✅ Usuario desaparece del listado
- ✅ Total de usuarios disminuye a 3
- ✅ Nota: Usuario se marca como inactivo en BD (soft delete)

---

### Test 6: Validaciones

**Test 6a: Nombre Vacío**
```
Acción: Intentar guardar sin ingresar nombre
Resultado: Mostrar error "El nombre no puede estar vacío"
```

**Test 6b: Nombre Muy Largo (>100 caracteres)**
```
Acción: Intentar guardar nombre con 101+ caracteres
Resultado: Mostrar error "El nombre no puede exceder 100 caracteres"
```

---

### Test 7: Verificar Base de Datos Directamente

```bash
# Conectar a la BD desde otro terminal
docker-compose exec db mysql -u usuariodb -p1234 proyecto_go

# Dentro de MySQL:
mysql> SELECT * FROM usuarios WHERE activo = 1;

# Resultado esperado: Lista de usuarios activos
# mysql> SELECT COUNT(*) FROM usuarios WHERE activo = 1;
# Debe coincidir con el listado web
```

---

### Test 8: Verificar Persistencia

**Acción:**
1. Agregar un nuevo usuario "Test Persistencia"
2. Detener contenedores: `docker-compose down`
3. Reiniciar: `docker-compose up -d`
4. Acceder nuevamente a listado

**Resultado Esperado:**
- ✅ Usuario "Test Persistencia" sigue presente
- ✅ Los datos persisten en volumen de MariaDB

---

## 🔍 Debugging

### Ver Logs del Servidor Web
```bash
docker-compose logs -f web
```

### Ver Logs de Base de Datos
```bash
docker-compose logs -f db
```

### Conectar a Base de Datos Interactivamente
```bash
docker-compose exec db mysql -u usuariodb -p1234 proyecto_go
```

### Inspeccionar Archivos dentro del Contenedor
```bash
docker-compose exec web ls -la /var/www/html
```

---

## 📊 Comparación de Tiempos

Prueba medir la velocidad en comparación con las versiones en microservicios:

**Monolítica (PHP):**
```bash
# Medir tiempo de respuesta
time curl http://localhost:8080/listar.php
```

**Microservicios (Go):**
```bash
# (Después de desplegar V2_Microservicio con Kubernetes)
time curl http://localhost:8081/api/usuarios
```

---

## 🛑 Detener la Aplicación

```bash
# Parar contenedores pero mantener datos
docker-compose stop

# Parar y eliminar contenedores (mantiene volúmenes)
docker-compose down

# Parar, eliminar todo incluyendo volúmenes
docker-compose down -v
```

---

## ✅ Checklist de Verificación

- [ ] Docker Compose up funciona sin errores
- [ ] Página principal carga correctamente
- [ ] Se muestran 3 usuarios de prueba
- [ ] Crear usuario funciona
- [ ] Editar usuario funciona
- [ ] Eliminar usuario funciona
- [ ] Los datos persisten tras restart
- [ ] No hay errores SQL en logs
- [ ] La conexión a BD es estable

---

## 🎯 Observaciones Importantes

1. **Arquitectura Monolítica:** Todo está acoplado en un mismo lugar
2. **Escalabilidad:** Para más usuarios/peticiones, necesitarías más memoria en el mismo servidor
3. **Actualización:** Cambiar cualquier cosa requiere reiniciar toda la aplicación
4. **Comparación:** Revisa V2_Microservicio/ para ver cómo esto mejora con microservicios

---

**Fecha:** 27 de Diciembre, 2025  
**Estado:** ✅ Lista para pruebas
