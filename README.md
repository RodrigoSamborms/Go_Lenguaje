# 🚀 Go Lenguaje - Evolución de Arquitectura Monolítica a Microservicios

**Proyecto educativo de migración arquitectónica: Apache + PHP + MariaDB → Go Microservicios + Kubernetes**

---

## 📌 Descripción General

Este repositorio documenta un **viaje de transformación arquitectónica** que demuestra cómo evolucionar desde una arquitectura web tradicional monolítica hacia una moderna arquitectura de microservicios containerizada y orquestada con Kubernetes.

### Evolución Arquitectónica

```
FASE 1: Arquitectura Monolítica          FASE 2: Microservicios                FASE 3: Distribuido (Futuro)
─────────────────────────────            ──────────────────────────            ────────────────────────────
                                         
┌─────────────────────────────┐          ┌──────────────┐
│   Apache Web Server         │          │  Frontend    │ (Go + HTTP)
│  ├─ Página HTML estática    │          │  :8080       │
│  ├─ Scripts PHP             │          └──────────────┘
│  └─ API directa a BD        │                  ↓
└─────────────────────────────┘          ┌──────────────┐
         ↓                                │  API Backend │ (Go + HTTP)
┌─────────────────────────────┐          │  :8081       │                ┌──────────────┐
│  MariaDB                    │          └──────────────┘                │  Master K8s  │
│  - Usuarios                 │                  ↓                       │  (x86)       │
│  - Datos de aplicación      │          ┌──────────────┐                └──────────────┘
│  - Lógica de negocio        │          │  Database    │                       ↓
└─────────────────────────────┘          │  MariaDB     │          ┌─────────────────────┐
                                         │  :3306       │          │  Worker Nodes       │
                                         └──────────────┘          │  ├─ x86 (Potente)   │
                                                                   │  ├─ RPi 1 (ARM v7)  │
                                         Orquestado por:           │  ├─ RPi 2 (ARM v7)  │
                                         Kubernetes               │  └─ RPi 3 (ARM v8)  │
                                                                   └─────────────────────┘
```

---

## 📁 Estructura del Repositorio

```
Go_Lenguaje/
│
├── README.md                          # Este archivo
├── Instalacion.md                     # Guía de instalación
│
├── Practicas/                         # Ejercicios básicos de Go
│   ├── hello-world.go
│   ├── variables.go
│   ├── constantes.go
│   ├── arrays.go
│   ├── slices.go
│   ├── maps.go
│   ├── for.go
│   ├── if.go
│   └── switch.go
│
├── Microservicio/                     # FASE 1: Versión inicial (Docker)
│   ├── docker-compose.yml
│   ├── Dockerfile
│   ├── main.go
│   ├── db.go
│   ├── altas.go
│   ├── cambios.go
│   ├── listas.go
│   ├── init.sql
│   ├── index.html
│   └── styles.css
│
└── V2_Microservicio/                  # FASE 2: Versión con Kubernetes ⭐
    ├── KUBERNETES_FIXES.md            # Documentación de fixes aplicados
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── init.sql
    │
    ├── backend/                       # API Backend
    │   ├── Dockerfile
    │   ├── main_api.go
    │   ├── db.go
    │   ├── altas.go
    │   ├── cambios.go
    │   └── listas.go
    │
    ├── frontend/                      # Frontend Web
    │   ├── Dockerfile
    │   ├── main_web.go
    │   ├── index.html
    │   ├── styles.css
    │   ├── altas_web.go
    │   ├── cambios_web.go
    │   └── listas_web.go
    │
    ├── k8s-db-deployment.yaml         # Deployment: Base de Datos
    ├── k8s-api-deployment.yaml        # Deployment: API Backend
    ├── k8s-web-deployment.yaml        # Deployment: Frontend
    ├── k8s-secrets.yaml               # Secrets: Credenciales
    ├── k8s-persistence.yaml           # Persistence: Volúmenes
    └── k8s-db-claim*.yaml             # PVC: Persistencia
```

---

## 🎯 Fases del Proyecto

### ✅ FASE 1: Arquitectura Monolítica (Referencia)
**Ubicación:** `Microservicio/`

**Características:**
- Apache Web Server
- PHP scripts
- MariaDB
- Aplicación Web monolítica

**Propósito:** Punto de partida para demostrar la arquitectura tradicional.

---

### ✅ FASE 2: Microservicios con Kubernetes (Actual - ⭐ Enfoque Principal)
**Ubicación:** `V2_Microservicio/`

**Arquitectura Implementada:**

```
┌─────────────────────────────────────────────────────┐
│            Kubernetes Cluster (Minikube)            │
├─────────────────────────────────────────────────────┤
│                                                     │
│  ┌──────────────┐    ┌──────────────┐             │
│  │   Frontend   │    │     API      │    ┌────┐  │
│  │  (Port 8080) │───→│  (Port 8081) │───→│ DB │  │
│  │  go-frontend │    │ go-backend   │    │    │  │
│  └──────────────┘    └──────────────┘    └────┘  │
│       Service             Service          Service│
│      NodePort          ClusterIP         ClusterIP│
│                                                    │
└─────────────────────────────────────────────────────┘
         ↓
    ┌─────────────┐
    │   Acceso    │
    │localhost:8080
    │(Port-Forward)
    └─────────────┘
```

**Tecnologías:**
- **Go 1.23** - Lenguaje de programación
- **Docker** - Containerización
- **Kubernetes** - Orquestación
- **MariaDB 10.11** - Base de datos relacional
- **HTTP/REST** - Comunicación entre servicios

**Características Implementadas:**
- ✅ Separación clara de responsabilidades (Frontend/API/DB)
- ✅ Desacoplamiento de servicios
- ✅ API REST agnóstica (consumible por cualquier cliente)
- ✅ Escalabilidad horizontal (réplicas de pods)
- ✅ Persistencia de datos con volúmenes
- ✅ Configuración externalizada (ConfigMaps, Secrets)
- ✅ Health checks y reintentos automáticos
- ✅ Actualización de imágenes sin downtime

**Operaciones CRUD Funcionales:**
- ✅ Listar usuarios
- ✅ Crear usuarios
- ✅ Editar usuarios
- ✅ Eliminar usuarios

---

### 🔮 FASE 3: Kubernetes Distribuido (Futuro)
**Objetivo:** Escalar a cluster de múltiples nodos

**Descripción:**
Deploying el mismo cluster K8s en múltiples máquinas:

```
┌─────────────────────────────────────────────────────────┐
│  Master (x86 - Máquina potente)                        │
│  ├─ Control Plane de Kubernetes                        │
│  └─ Base de Datos (preferentemente aquí)               │
└─────────────────────────────────────────────────────────┘
              ↓
    ┌─────────────────┬──────────────┬──────────────┐
    ↓                 ↓              ↓              ↓
┌──────────────┐ ┌──────────────┐ ┌──────────────┐
│ Worker (x86) │ │ Worker RPi 1  │ │ Worker RPi 2 │
│   (arm64)    │ │   (armv7)     │ │   (armv8)    │
└──────────────┘ └──────────────┘ └──────────────┘
    Frontend       API            Frontend
    API            Frontend       API
```

**Beneficios:**
- 🔄 Alta disponibilidad
- 📈 Escalado automático
- 💰 Optimización de costos (Raspberry Pi es económica)
- 🛡️ Redundancia de datos
- ⚡ Balance de carga automático

---

## 🚀 Cómo Ejecutar FASE 2 (Actual)

### Requisitos Previos
```bash
# WSL 2 o Linux
# Docker Desktop con Kubernetes habilitado
# O Minikube
minikube version
kubectl version --client
```

### Paso 1: Iniciar Minikube
```bash
minikube start --driver=docker
```

### Paso 2: Desplegar a Kubernetes
```bash
cd V2_Microservicio

# Aplicar todos los manifests
kubectl apply -f k8s-secrets.yaml
kubectl apply -f k8s-db-deployment.yaml
kubectl apply -f k8s-api-deployment.yaml
kubectl apply -f k8s-web-deployment.yaml
```

### Paso 3: Verificar Estado
```bash
kubectl get pods
kubectl get svc
```

### Paso 4: Acceder a la Aplicación

**Opción A: Port-Forward (Recomendado)**
```bash
kubectl port-forward svc/web-service 8080:8080
# Acceder: http://localhost:8080
```

**Opción B: NodePort directo**
```bash
minikube ip  # Obtener IP
kubectl get svc web-service  # Obtener puerto
# Ejemplo: http://192.168.58.2:30198
```

---

## 📚 Documentación Importante

- [KUBERNETES_FIXES.md](V2_Microservicio/KUBERNETES_FIXES.md) - Problemas resueltos y soluciones aplicadas
- [Instalacion.md](Instalacion.md) - Guía de instalación de herramientas
- [Practicas/Introduccion.md](Practicas/Introduccion.md) - Introducción a Go

---

## 🔑 Conceptos Clave Aprendidos

### Transición de Arquitecturas

| Aspecto | Monolítica | Microservicios |
|---------|-----------|-----------------|
| **Despliegue** | Todo junto | Servicios independientes |
| **Escalabilidad** | Vertical (más recursos) | Horizontal (más réplicas) |
| **Actualización** | Reiniciar toda la app | Actualizar servicio específico |
| **Tolerancia a fallos** | Falla todo | Solo falla un servicio |
| **Tecnologías** | Todas iguales | Heterogéneas (cada una la mejor) |
| **Complejidad** | Baja inicial | Alta pero escalable |

### Tecnologías Clave

- **Containerización:** Encapsular aplicaciones con Docker
- **Orquestación:** Gestionar containers con Kubernetes
- **Microservicios:** Separar por responsabilidad funcional
- **API REST:** Comunicación entre servicios
- **Variables de Entorno:** Configuración sin cambiar código
- **Volúmenes:** Persistencia de datos

---

## 💡 Ventajas de Esta Arquitectura

✅ **Escalabilidad:** Cada servicio puede escalarse independientemente  
✅ **Resiliencia:** Si un servicio falla, otros continúan  
✅ **Facilidad de desarrollo:** Teams pueden trabajar en paralelo  
✅ **Deployment flexible:** Actualizar sin downtime  
✅ **Portabilidad:** Funciona en cualquier infraestructura (x86, ARM, cloud)  
✅ **Eficiencia de recursos:** Usar solo los recursos necesarios  

---

## 🔜 Próximos Pasos Sugeridos

### Mejoras Inmediatas
- [ ] Agregar autenticación JWT
- [ ] Implementar validación de datos
- [ ] Agregar logging centralizado
- [ ] Configurar monitoring (Prometheus)
- [ ] Agregar health checks

### Para Fase 3 (Distribuida)
- [ ] Preparar imágenes multi-arquitectura (amd64, arm7, arm8)
- [ ] Configurar ingress controller
- [ ] Implementar persistent storage distribuido (etcd)
- [ ] Configurar networking entre nodos

---

## 📖 Recursos Útiles

- [Documentación oficial de Kubernetes](https://kubernetes.io/docs/)
- [Documentación oficial de Go](https://golang.org/doc/)
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [Minikube](https://minikube.sigs.k8s.io/)

---

## 👤 Autor

**Sambo** - Proyecto de aprendizaje y demostración arquitectónica

---

## 📝 Licencia

Proyecto educativo - Libre para uso personal y educativo

---

## 🎓 Conclusión

Este proyecto demuestra cómo evolucionar desde una arquitectura tradicional monolítica hacia una moderna arquitectura de microservicios, manteniendo la funcionalidad mientras se ganan beneficios de escalabilidad, resiliencia y mantenibilidad.

La transición no es solo sobre tecnología, es sobre **cambiar la forma de pensar** en cómo diseñamos, deployamos y escalamos aplicaciones.

**¡Próximo objetivo:** Distribuir este cluster en múltiples nodos (x86 + Raspberry Pi) con Kubernetes! 🚀

---

**Estado del Proyecto:** ✅ OPERACIONAL  
**Última actualización:** 27 de Diciembre, 2025
