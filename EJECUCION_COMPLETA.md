# 🚀 Guía de Ejecución - Todas las Arquitecturas

Esta guía te permite ejecutar y comparar las **3 versiones** de la aplicación de forma simultánea.

---

## 📋 Requisitos Previos

```bash
# Verificar que tengas instalado:
docker --version        # Docker Desktop
docker-compose --version # Docker Compose
kubectl version --client # Kubectl (para Kubernetes)
minikube version        # Minikube (para Kubernetes)
```

---

## 🎯 Ejecución en Paralelo

### Terminal 1: Monolítica (PHP)
```bash
cd Go_Lenguaje/Monolotico
docker-compose up -d

# Acceder: http://localhost:8080
# Ver logs: docker-compose logs -f web
```

### Terminal 2: Microservicio v1 (Docker Compose)
```bash
cd Go_Lenguaje/Microservicio
docker-compose up -d

# Frontend: http://localhost:8080
# API: http://localhost:8081/api/usuarios

# Problema: Puerto 8080 duplicado
# Solución: cambiar puerto en docker-compose.yml
# ports:
#   - "8081:8080"  <- cambiar primer puerto
```

### Terminal 3: Microservicio v2 (Kubernetes)
```bash
cd Go_Lenguaje/V2_Microservicio

# Iniciar Minikube
minikube start

# Aplicar manifests
kubectl apply -f k8s-secrets.yaml
kubectl apply -f k8s-db-deployment.yaml
kubectl apply -f k8s-api-deployment.yaml
kubectl apply -f k8s-web-deployment.yaml

# Port-forward (en otra terminal)
kubectl port-forward svc/web-service 8082:8080

# Acceder: http://localhost:8082
```

---

## 🔄 Gestión de Puertos

Para evitar conflictos de puertos, usa estos:

| Versión | Puerto | URL |
|---------|--------|-----|
| Monolítica (PHP) | 8080 | http://localhost:8080 |
| Microservicio v1 (Frontend) | 8081 | http://localhost:8081 |
| Microservicio v1 (API) | 8082 | http://localhost:8082 |
| Microservicio v2 (Port-forward) | 8083 | http://localhost:8083 |

**Modificar docker-compose.yml para v1:**
```yaml
# Cambiar ports en docker-compose.yml
web-service:
  ports:
    - "8081:8080"  # ← Cambiar de 8080

api-service:
  ports:
    - "8082:8081"  # ← Cambiar de 8081
```

---

## 📊 Comparación en Tiempo Real

### 1. Probar Velocidad

```bash
# Terminal separada - Medir tiempos
echo "=== Monolítica ==="
time curl -s http://localhost:8080/listar.php > /dev/null

echo "=== Microservicio v1 ==="
time curl -s http://localhost:8081/listar > /dev/null

echo "=== Microservicio v2 ==="
time curl -s http://localhost:8083/listar > /dev/null
```

### 2. Verificar Estado

```bash
# Monolítica
docker-compose -f Monolotico/docker-compose.yml ps

# Microservicio v1
docker-compose -f Microservicio/docker-compose.yml ps

# Microservicio v2
kubectl get pods
kubectl get svc
```

### 3. Ver Logs

```bash
# Monolítica
docker-compose -f Monolotico/docker-compose.yml logs -f web

# Microservicio v1
docker-compose -f Microservicio/docker-compose.yml logs -f web

# Microservicio v2
kubectl logs -f deployment/web-service
```

---

## 🧪 Test Comparativo

### Crear Usuario en Cada Una

```bash
# Monolítica
curl -X POST http://localhost:8080/agregar.php \
  -d "nombre=Test+Monolitica"

# Microservicio v1
curl -X POST http://localhost:8081/agregar \
  -d "nombre=Test+Microservicio+v1"

# Microservicio v2
curl -X POST http://localhost:8083/agregar \
  -d "nombre=Test+Microservicio+v2"
```

### Listar Usuarios

```bash
# Todas devuelven HTML renderizado
curl http://localhost:8080/listar.php
curl http://localhost:8081/listar
curl http://localhost:8083/listar
```

---

## 🛑 Detener Todo

```bash
# Monolítica
cd Monolotico && docker-compose down

# Microservicio v1
cd Microservicio && docker-compose down

# Microservicio v2
kubectl delete -f V2_Microservicio/k8s-*.yaml
minikube stop
```

---

## 📈 Simulación de Escenarios

### Escenario 1: Falla del Servicio Web

#### Monolítica
```bash
# Detener contenedor web
docker-compose -f Monolotico/docker-compose.yml stop web
# Resultado: APLICACIÓN COMPLETA CAÍDA ❌
# Reiniciar: docker-compose -f Monolotico/docker-compose.yml up -d web
```

#### Microservicio v1
```bash
# Detener frontend
docker-compose -f Microservicio/docker-compose.yml stop web
# Resultado: Frontend caído, API sigue activa ⚠️
# Verificar: curl http://localhost:8082/api/usuarios (sigue funcionando)
```

#### Microservicio v2
```bash
# Eliminar pod del frontend
kubectl delete pod -l app=web
# Resultado: K8s automáticamente recrea el pod ✅
# Tiempo de recuperación: ~10 segundos
```

---

### Escenario 2: Falla de Base de Datos

#### Monolítica
```bash
docker-compose -f Monolotico/docker-compose.yml stop db
# Resultado: TODA LA APLICACIÓN FALLA ❌
# Tiempo para detectar falla: Inmediato
```

#### Microservicio v1
```bash
docker-compose -f Microservicio/docker-compose.yml stop db
# Resultado: API falla al intentar conectarse ⚠️
# Frontend sigue sirviendo HTML (pero sin datos)
# Los logs muestran errores de conexión
```

#### Microservicio v2
```bash
kubectl delete pod -l app=db
# Resultado: K8s recrea el pod automáticamente ✅
# Con persistencia: Los datos se recuperan
# Tiempo: ~15-20 segundos
```

---

### Escenario 3: Escalado de Carga

#### Monolítica
```bash
# Única opción: detener todo y subir a un servidor más grande ❌
# Downtime: Sí
# Complejidad: Alta
```

#### Microservicio v1
```bash
# Escalar manualmente:
# 1. Abrir docker-compose.yml
# 2. Aumentar replicas
# 3. Reconstruir imágenes
# 4. Levantar de nuevo

# Complejidad: Moderada
# Downtime: Parcial
```

#### Microservicio v2
```bash
# Escalar automáticamente:
kubectl scale deployment web-service --replicas=3
kubectl scale deployment api-service --replicas=3

# Complejidad: Baja
# Downtime: No
# Tiempo: ~20 segundos
```

---

## 📊 Dashboard de Monitoreo

### Para Microservicio v2 (Kubernetes)

```bash
# Ver recursos en tiempo real
watch kubectl top pods

# Ver eventos del cluster
watch kubectl get events --sort-by='.lastTimestamp'

# Dashboard de Minikube
minikube dashboard
```

---

## 🔍 Inspección de Datos

### Verificar Datos en BD

```bash
# Monolítica
docker-compose -f Monolotico/docker-compose.yml exec db \
  mysql -u usuariodb -p1234 proyecto_go \
  -e "SELECT * FROM usuarios WHERE activo = 1;"

# Microservicio v1
docker-compose -f Microservicio/docker-compose.yml exec db \
  mysql -u usuariodb -p1234 proyecto_go \
  -e "SELECT * FROM usuarios WHERE activo = 1;"

# Microservicio v2
kubectl exec -it $(kubectl get pod -l app=db -o jsonpath='{.items[0].metadata.name}') \
  -- mysql -u usuariodb -p1234 proyecto_go \
  -e "SELECT * FROM usuarios WHERE activo = 1;"
```

---

## 📝 Registro de Observaciones

Mientras ejecutas las 3 versiones, toma nota de:

- [ ] Tiempo de startup
- [ ] Tiempo de respuesta promedio
- [ ] Consumo de recursos (CPU, memoria)
- [ ] Facilidad de escalado
- [ ] Recuperación ante fallos
- [ ] Facilidad de actualización
- [ ] Mantenimiento requerido

---

## 🎓 Conclusiones Esperadas

Al completar este ejercicio, deberías observar:

1. **Monolítica es simple pero frágil**
   - Fácil de entender
   - Difícil de recuperarse de fallos

2. **Microservicio v1 es mejor pero aún limitado**
   - Mejor separación de responsabilidades
   - Aún requiere gestión manual

3. **Microservicio v2 es robusto y automático**
   - Auto-recuperación
   - Escalado automático
   - Mayor complejidad inicial

---

## 🆘 Solución de Problemas

### Puertos en Uso

```bash
# Encontrar qué está usando puerto 8080
lsof -i :8080  # macOS/Linux
netstat -ano | findstr :8080  # Windows
```

### Contenedores No Inician

```bash
# Ver logs de error
docker-compose logs web
docker-compose logs db

# Reintentar con limpieza
docker-compose down -v
docker-compose up -d
```

### Kubernetes No Responde

```bash
# Reiniciar Minikube
minikube stop
minikube start --driver=docker

# Limpiar recursos
kubectl delete all --all
```

---

**Última actualización:** 27 de Diciembre, 2025  
**Estado:** ✅ Listo para ejecutar
