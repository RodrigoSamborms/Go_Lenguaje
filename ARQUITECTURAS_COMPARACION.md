# 📚 Documentación Comparativa: Arquitecturas

## 🎯 Resumen Ejecutivo

Este documento compara las **tres versiones** de la misma aplicación CRUD en diferentes arquitecturas.

---

## 📊 Tabla Comparativa General

| Característica | Monolítica (PHP) | Microservicio v1 (Go) | Microservicio v2 (Go + K8s) |
|---|---|---|---|
| **Carpeta** | `Monolotico/` | `Microservicio/` | `V2_Microservicio/` |
| **Lenguaje** | PHP 8.2 | Go 1.23 | Go 1.23 |
| **Servidor Web** | Apache | Built-in Go | Built-in Go |
| **Base de Datos** | MariaDB integrada | MariaDB integrada | MariaDB separada |
| **Orquestación** | Docker Compose | Docker Compose | Kubernetes |
| **API** | ❌ No (acoplada) | ✅ REST (desacoplada) | ✅ REST (desacoplada) |
| **Escalabilidad** | Vertical | Horizontal (limitada) | Horizontal (completa) |
| **Frontend** | PHP en servidor | Separado en Go | Servicio independiente |
| **Tolerancia fallos** | ❌ Falla todo | ⚠️ Parcial | ✅ Aislada por servicio |
| **Actualización** | ⚠️ Downtime total | ⚠️ Downtime parcial | ✅ Sin downtime |
| **Facilidad uso** | ✅ Simple | ⚠️ Moderada | ⚠️ Compleja |

---

## 🏗️ Estructura Arquitectónica

### Monolítica (PHP)
```
┌─────────────────────────────────┐
│   Apache + PHP (Puerto 8080)    │
├─────────────────────────────────┤
│  ├─ index.php (Página)          │
│  ├─ listar.php (Lógica)         │
│  ├─ agregar.php (Lógica)        │
│  ├─ editar.php (Lógica)         │
│  └─ borrar.php (Lógica)         │
└─────────────────────────────────┘
          ↓
┌─────────────────────────────────┐
│  MariaDB (Puerto 3306)          │
│  ├─ proyecto_go (BD)            │
│  └─ usuarios (Tabla)            │
└─────────────────────────────────┘

Orquestación: Docker Compose
Escalabilidad: Vertical
```

### Microservicios v1 (Go + Docker Compose)
```
┌──────────────────────────┐
│ Frontend Web (Go)        │
│ Puerto 8080              │
├──────────────────────────┤
│  Sirve HTML estático     │
│  Hace requests HTTP      │
└──────────────────────────┘
         ↓ HTTP
┌──────────────────────────┐
│ Backend API (Go)         │
│ Puerto 8081              │
├──────────────────────────┤
│  GET /api/usuarios       │
│  POST /api/usuarios/*    │
└──────────────────────────┘
         ↓ SQL
┌──────────────────────────┐
│ MariaDB (Puerto 3306)    │
├──────────────────────────┤
│  proyecto_go (BD)        │
│  usuarios (Tabla)        │
└──────────────────────────┘

Orquestación: Docker Compose
Escalabilidad: Horizontal (limitada a 1 máquina)
```

### Microservicios v2 (Go + Kubernetes)
```
┌─────────────────────────────────────────┐
│        Kubernetes Cluster               │
├─────────────────────────────────────────┤
│                                         │
│  ┌──────────────┐  ┌──────────────┐   │
│  │  Frontend    │  │  Backend API │   │
│  │  (Servicio)  │→ │  (Servicio)  │   │
│  │  Puerto 8080 │  │  Puerto 8081 │   │
│  └──────────────┘  └──────────────┘   │
│                            ↓           │
│                    ┌──────────────┐   │
│                    │  MariaDB     │   │
│                    │  (Servicio)  │   │
│                    │  Puerto 3306 │   │
│                    └──────────────┘   │
│                                         │
└─────────────────────────────────────────┘

Orquestación: Kubernetes
Escalabilidad: Horizontal (distribuido)
Tolerancia fallos: Por servicio
```

---

## 💻 Cómo Ejecutar Cada Una

### Monolítica (PHP)

**Inicio:**
```bash
cd Monolotico/
docker-compose up -d
# Acceder: http://localhost:8080
```

**Detener:**
```bash
docker-compose down
```

---

### Microservicio v1 (Docker Compose)

**Inicio:**
```bash
cd Microservicio/
docker-compose up -d
# Frontend: http://localhost:8080
# API: http://localhost:8081/api/usuarios
```

**Detener:**
```bash
docker-compose down
```

---

### Microservicio v2 (Kubernetes)

**Requisitos:**
- Minikube o Docker Desktop con K8s habilitado

**Inicio:**
```bash
cd V2_Microservicio/

# Aplicar manifests
kubectl apply -f k8s-secrets.yaml
kubectl apply -f k8s-db-deployment.yaml
kubectl apply -f k8s-api-deployment.yaml
kubectl apply -f k8s-web-deployment.yaml

# Port-forward
kubectl port-forward svc/web-service 8080:8080
# Acceder: http://localhost:8080
```

**Detener:**
```bash
kubectl delete -f k8s-*.yaml
```

---

## 📈 Comparación de Funcionalidades

### CRUD Implementado

Todas las tres versiones implementan las **mismas 4 operaciones:**

| Operación | Monolítica | Microservicio v1 | Microservicio v2 |
|-----------|-----------|------------------|------------------|
| **Listar** | ✅ listar.php | ✅ GET /api/usuarios | ✅ GET /api/usuarios |
| **Crear** | ✅ agregar.php | ✅ POST /api/usuarios/crear | ✅ POST /api/usuarios/crear |
| **Editar** | ✅ editar.php | ✅ POST /api/usuarios/editar | ✅ POST /api/usuarios/editar |
| **Eliminar** | ✅ borrar.php | ✅ POST /api/usuarios/borrar | ✅ POST /api/usuarios/borrar |

---

## 🔄 Flujo de Petición

### Monolítica
```
Usuario → Navegador → Apache/PHP → Ejecuta lógica → MariaDB
                                 ↓
                            Retorna HTML
```

### Microservicio v1
```
Usuario → Navegador → Web (Go) → Solicita JSON → API (Go) → MariaDB
                            ↑                            ↓
                            ← Recibe JSON ←
         ↓
    Renderiza HTML
```

### Microservicio v2
```
Usuario → Navegador → Web (K8s) → Solicita JSON → API (K8s) → MariaDB (K8s)
                            ↑                           ↓
                            ← Recibe JSON ←
         ↓
    Renderiza HTML
```

---

## 🚀 Escalabilidad

### Monolítica
```
Demanda ↑ → Necesitas Servidor más Grande (Vertical)
           ❌ Caro
           ❌ Requiere downtime
           ❌ Cuello de botella en BD
```

### Microservicio v1
```
Demanda ↑ → Replica contenedores (Horizontal)
          ⚠️ Limitado a 1 máquina
          ⚠️ Orquestación manual
          ⚠️ Gestión compleja de networking
```

### Microservicio v2
```
Demanda ↑ → Auto-scale horizontal (K8s)
           ✅ Distribuido en múltiples máquinas
           ✅ Orquestación automática
           ✅ Tolera fallos de máquinas
           ✅ Load balancing automático
```

---

## 🔧 Gestión de Actualizaciones

### Monolítica
```
Cambio en código ↓
Reconstruir imagen ↓
docker-compose down (❌ DOWNTIME)
docker-compose up
```

### Microservicio v1
```
Cambio en API ↓
Reconstruir imagen de backend ↓
docker-compose up -d backend (⚠️ DOWNTIME solo API)
```

### Microservicio v2
```
Cambio en API ↓
Reconstruir imagen ↓
kubectl rollout restart deployment api-service (✅ SIN DOWNTIME)
```

---

## 💰 Costo de Infraestructura

### Monolítica
```
1 Servidor Potente = $$$$
Difícil de underutilizar
```

### Microservicio v1
```
3 Contenedores en 1 máquina = $$$
Mejor utilización que monolítica
```

### Microservicio v2
```
1 x86 (Master) + 3 Raspberry Pi (Workers) = $$
Escalable económicamente
Distribuido en hardware heterogéneo
```

---

## 📚 Documentación

| Versión | Ubicación | Documentación |
|---------|-----------|---------------|
| Monolítica | `Monolotico/` | [README.md](Monolotico/README.md), [TESTING.md](Monolotico/TESTING.md) |
| Microservicio v1 | `Microservicio/` | README.md (en repo) |
| Microservicio v2 | `V2_Microservicio/` | [KUBERNETES_FIXES.md](V2_Microservicio/KUBERNETES_FIXES.md) |
| Visión General | `./` | [README.md](README.md) |

---

## 🎓 Lecciones Clave

### 1. Monolítica
- ✅ Fácil de entender para principiantes
- ✅ Rápida para prototipado
- ❌ Difícil de escalar
- ❌ Cambios afectan toda la aplicación

### 2. Microservicios v1
- ✅ Separación de responsabilidades
- ✅ API desacoplada
- ✅ Escalabilidad horizontal básica
- ❌ Limitado a 1 máquina
- ❌ Complejidad aumenta

### 3. Microservicios v2
- ✅ Verdadera arquitectura distribuida
- ✅ Alta disponibilidad
- ✅ Auto-healing
- ✅ Escalado automático
- ❌ Complejidad operacional
- ❌ Requiere más herramientas

---

## 🔜 Progresión Natural

```
Proyecto Pequeño
    ↓
Monolítica (PHP) ← AQUÍ ESTAMOS
    ↓
    (Aplicación crece)
    ↓
Microservicio v1 (Docker Compose)
    ↓
    (Necesidad de HA, múltiples máquinas)
    ↓
Microservicio v2 (Kubernetes)
    ↓
    (Distribuido en RPi + x86)
    ↓
Fase 3: Cluster Distribuido
```

---

## 🧪 Ejercicio Propuesto

1. **Ejecutar las 3 versiones localmente**
   - Monolítica en puerto 8080
   - v1 en puerto 8081
   - v2 con port-forward en 8082 (o 8080 con port-forward diferente)

2. **Comparar velocidades**
   ```bash
   time curl http://localhost:8080/listar  # Monolítica
   time curl http://localhost:8081/api/usuarios  # v1
   time curl http://localhost:8082/listar  # v2
   ```

3. **Simular falla**
   - Monolítica: Detener contenedor → Todo falla
   - v1: Detener API → Frontend no funciona
   - v2: Detener un pod → K8s recrea automáticamente

4. **Medir escalabilidad**
   - Intentar escalar cada una
   - Observar cómo la facilidad aumenta

---

## 📝 Conclusión

Estas **tres versiones de la misma aplicación** demuestran la **evolución natural** de:
- Monolítica → Microservicios → Arquitectura Distribuida

Cada una tiene su lugar según el contexto:
- **Pequeño prototipo:** Monolítica
- **Aplicación en crecimiento:** Microservicios v1
- **Producción de alto rendimiento:** Microservicios v2

---

**Creado:** 27 de Diciembre, 2025  
**Estado:** ✅ Completo
