# Documentación de Fixes para Kubernetes - V2 Microservicio

**Fecha:** 27 de Diciembre, 2025  
**Estado:** ✅ Completado y Funcional

---

## 📋 Resumen Ejecutivo

Se identificaron y solucionaron **3 problemas críticos** que impedían que la aplicación funcionara correctamente en Kubernetes:

1. **Frontend hardcodeado** - No leía variables de entorno
2. **URL incorrecta del API** - Apuntaba a localhost en lugar del servicio K8s
3. **Conflicto de variables de entorno** - Kubernetes sobrescribía `DB_PORT` automáticamente

---

## 🐛 Problemas Identificados

### Problema 1: Frontend No Leía Variable de Entorno `API_URL`

**Ubicación:** `frontend/main_web.go`

**Descripción:**
El código frontend tenía la URL del API **hardcodeada** como una constante:
```go
const apiURL = "http://api-service:8081/api/usuarios" //ejecución en Docker
```

Aunque el deployment de Kubernetes definía la variable `API_URL`, el código no la estaba usando.

**Impacto:** El frontend intentaba conectarse a direcciones incorrectas, resultando en errores de conexión rechazada.

---

### Problema 2: Configuración Incorrecta en k8s-web-deployment.yaml

**Ubicación:** `k8s-web-deployment.yaml`

**Descripción:**
La variable de entorno `API_URL` estaba mal configurada:
```yaml
env:
- name: API_URL
  #value: "http://api-service:8081/api/usuarios"
  value: "http://127.0.0.1:33587/api/usuarios"  # ❌ INCORRECTA
```

Se intentaba conectar a `localhost` en lugar del servicio Kubernetes `api-service`.

**Impacto:** Incluso si el frontend leyera la variable, apuntaría a una dirección interna incorrecta.

---

### Problema 3: Conflicto de Variables de Entorno en K8s

**Ubicación:** `k8s-api-deployment.yaml`

**Descripción:**
Kubernetes inyecta automáticamente variables de entorno para cada servicio. Para el servicio `db`, crea:
- `DB_PORT=tcp://10.108.220.151:3306` (URL completa)
- `DB_SERVICE_HOST=10.108.220.151`
- `DB_SERVICE_PORT=3306`

El código Go leía `DB_PORT` esperando solo el número del puerto (`"3306"`), pero recibía una URL completa (`"tcp://10.108.220.151:3306"`), lo que causaba errores de conexión.

**Impacto:** El API no podía conectarse a la base de datos, causando `CrashLoopBackOff`.

**Error en logs:**
```
Intentando conectar a: usuariodb@tcp(db:tcp://10.108.220.151:3306)/proyecto_go
Error: dial tcp: lookup db:tcp://10.108.220.151:3306: no such host
```

---

## ✅ Soluciones Implementadas

### Solución 1: Modificar Frontend para Leer Variable de Entorno

**Archivo:** `frontend/main_web.go`

**Cambios:**
```go
// ANTES (Constante hardcodeada)
const apiURL = "http://api-service:8081/api/usuarios"

// DESPUÉS (Lee variable de entorno con fallback)
func getAPIURL() string {
	if url := os.Getenv("API_URL"); url != "" {
		return url
	}
	// Fallback para ejecución en Docker/Kubernetes
	return "http://api-service:8081/api/usuarios"
}

var apiURL = getAPIURL()
```

**Beneficio:** El frontend ahora puede usar diferentes URLs según donde se ejecute (local, Docker, K8s).

---

### Solución 2: Actualizar k8s-web-deployment.yaml

**Archivo:** `k8s-web-deployment.yaml`

**Cambios:**
```yaml
# ANTES (URL incorrecta)
env:
- name: API_URL
  value: "http://127.0.0.1:33587/api/usuarios"

# DESPUÉS (URL correcta del servicio K8s)
env:
- name: API_URL
  value: "http://api-service:8081/api/usuarios"
```

**Beneficio:** El frontend conecta correctamente al servicio `api-service` dentro del cluster.

---

### Solución 3: Configurar DB_PORT Explícitamente en k8s-api-deployment.yaml

**Archivo:** `k8s-api-deployment.yaml`

**Cambios:**
```yaml
# ANTES (Sin especificar DB_PORT, K8s lo sobrescribía)
env:
- name: DB_HOST
  value: "db"

# DESPUÉS (DB_PORT explícito sobrescribe la variable automática)
env:
- name: DB_HOST
  value: "db"
- name: DB_PORT
  value: "3306"  # Puerto explícito para sobrescribir variable automática de K8s
- name: DB_USER
  valueFrom:
    configMapKeyRef:
      name: db-config
      key: mysql-user
- name: DB_PASS
  valueFrom:
    secretKeyRef:
      name: db-secrets
      key: mysql-password
- name: DB_NAME
  valueFrom:
    configMapKeyRef:
      name: db-config
      key: mysql-database
```

También se agregó `imagePullPolicy: Always` para forzar la descarga de la última versión de las imágenes:
```yaml
containers:
- name: api-container
  image: samborms/go-backend:v2
  imagePullPolicy: Always  # Forzar descarga de la última versión
```

**Beneficio:** Las variables de entorno se configuran correctamente, permitiendo la conexión exitosa a la base de datos.

---

### Solución 4: Agregar Logs de Depuración en backend/db.go

**Archivo:** `backend/db.go`

**Cambios:**
Se agregaron logs informativos para facilitar la depuración:
```go
// Log de configuración (sin mostrar la contraseña completa)
fmt.Printf("Intentando conectar a: %s@tcp(%s:%s)/%s\n", user, host, port, name)

// En el bucle de reintento
if err != nil {
    fmt.Printf("Error: %v\n", err)
}
```

**Beneficio:** Facilita la identificación rápida de problemas de conexión en futuros despliegues.

---

## 🔧 Archivos Modificados

| Archivo | Cambios |
|---------|---------|
| `frontend/main_web.go` | Agregada función `getAPIURL()` para leer variable de entorno |
| `k8s-web-deployment.yaml` | Corregida URL del API de `127.0.0.1:33587` a `api-service:8081` |
| `k8s-api-deployment.yaml` | Agregados `DB_PORT` explícito e `imagePullPolicy: Always` |
| `backend/db.go` | Agregados logs de depuración para facilitar troubleshooting |

---

## 📦 Imágenes Docker Reconstruidas y Desplegadas

### Backend
```bash
docker build -f backend/Dockerfile -t samborms/go-backend:v2 .
docker push samborms/go-backend:v2
```

### Frontend
```bash
docker build -f frontend/Dockerfile -t samborms/go-frontend:v2 .
docker push samborms/go-frontend:v2
```

---

## 🚀 Cómo Acceder a la Aplicación

### Opción 1: Mediante Minikube IP + NodePort (Directo)

1. **Obtener IP de Minikube:**
   ```bash
   minikube ip
   ```
   Resultado: `192.168.58.2`

2. **Obtener puerto NodePort:**
   ```bash
   kubectl get svc web-service
   ```
   Resultado: Puerto `30198` (puede variar)

3. **Acceder a la aplicación:**
   ```
   http://192.168.58.2:30198
   ```

### Opción 2: Port-Forward (Recomendado - Usa localhost)

Esta opción redirige el puerto del servicio K8s a tu máquina local:

1. **Ejecutar port-forward:**
   ```bash
   kubectl port-forward svc/web-service 8080:8080
   ```

2. **Resultado esperado:**
   ```
   Forwarding from 127.0.0.1:8080 -> 8080
   Forwarding from [::1]:8080 -> 8080
   Listening on port 8080.
   ```

3. **Acceder a la aplicación en tu navegador:**
   ```
   http://localhost:8080
   ```

**Ventajas del Port-Forward:**
- ✅ Acceso directo en `localhost` (más simple)
- ✅ No necesita conocer puertos dinámicos
- ✅ Seguridad mejorada (solo acceso local)
- ✅ Perfecto para desarrollo local

**Para detener el port-forward:**
Presiona `Ctrl + C` en la terminal donde ejecutaste el comando.

### Verificar Estado de los Pods

```bash
# Ver estado de todos los pods
kubectl get pods

# Resultado esperado:
# NAME                          READY   STATUS    RESTARTS   AGE
# api-service-58766756c5-h9tk8   1/1    Running   0          2m
# db-57c7b76c4f-w2spd            1/1    Running   0          90m
# web-service-d954c5f4-jwwcz     1/1    Running   0          10m
```

### Ver Logs

```bash
# Logs del API
kubectl logs -f deployment/api-service

# Logs del Frontend
kubectl logs -f deployment/web-service

# Logs de la Base de Datos
kubectl logs -f deployment/db
```

---

## ✨ Resultados Finales

### ✅ API Conectado a Base de Datos
```
Intentando conectar a: usuariodb@tcp(db:3306)/proyecto_go
¡Conexión exitosa a la base de datos!
Backend API corriendo en http://localhost:8081
```

### ✅ Frontend Accesible
```html
<!-- Respuesta exitosa con datos de la base de datos -->
<h1>Usuarios Activos (Desde Microservicio API)</h1>
<ul>
    <li>[1] Rodrigo</li>
    <li>[2] Sambo</li>
    <li>[3] Go Developer</li>
</ul>
```

### ✅ Operaciones CRUD Funcionales
- ✅ Listar usuarios
- ✅ Agregar nuevos usuarios
- ✅ Editar usuarios existentes
- ✅ Borrar usuarios

---

## 🔍 Lecciones Aprendidas

### 1. Variables de Entorno Automáticas en Kubernetes
Kubernetes inyecta automáticamente variables para cada servicio. Siempre especifica explícitamente las variables críticas para evitar conflictos.

### 2. Importancia de los Logs
Agregar logs detallados facilita enormemente la depuración en producción.

### 3. Diferencias Entre Entornos
Lo que funciona en Docker local puede fallar en Kubernetes debido a diferencias en:
- DNS interno
- Inyección de variables de entorno
- Políticas de red

### 4. imagePullPolicy
Usar `imagePullPolicy: Always` durante desarrollo asegura que siempre se use la última versión de la imagen.

---

## 📝 Comandos Útiles de Referencia

```bash
# Aplicar todos los manifests K8s
kubectl apply -f k8s-*.yaml

# Reiniciar un deployment
kubectl rollout restart deployment <nombre>

# Ver eventos del cluster
kubectl get events

# Describir un pod para ver detalles
kubectl describe pod <nombre-pod>

# Ejecutar comando dentro de un pod
kubectl exec -it <nombre-pod> -- <comando>

# Limpiar pods completados
kubectl delete pod --field-selector status.phase=Failed
```

---

## 🎯 Próximos Pasos (Opcional)

1. **Agregar health checks:** Implementar liveness y readiness probes
2. **Configurar auto-scaling:** Ajustar HPA basado en CPU/memoria
3. **Implementar ingress:** Exponer la aplicación sin NodePort
4. **Monitoring:** Agregar Prometheus/Grafana para monitoreo
5. **CI/CD:** Automatizar builds y deployments con GitLab CI o GitHub Actions

---

**Estado Final:** ✅ **COMPLETADO Y FUNCIONAL**  
**Última actualización:** 27 de Diciembre, 2025
