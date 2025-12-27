# 🗂️ ÍNDICE COMPLETO - Go Lenguaje Repository

**Última actualización:** 27 de Diciembre, 2025

---

## 📍 Navegación Rápida

### 🎯 Si eres Nuevo en el Repositorio
1. **Comienza aquí:** [README.md](README.md) - Descripción general del proyecto
2. **Entiende la evolución:** [ARQUITECTURAS_COMPARACION.md](ARQUITECTURAS_COMPARACION.md)
3. **Ejecuta las demos:** [EJECUCION_COMPLETA.md](EJECUCION_COMPLETA.md)

### 📚 Si quieres Aprender Go
1. **Ejercicios básicos:** [Practicas/Introduccion.md](Practicas/Introduccion.md)
2. **Instala herramientas:** [Instalacion.md](Instalacion.md)
3. **Explora ejemplos:** Carpeta [Practicas/](Practicas/)

### 🏗️ Si quieres Entender Arquitecturas
1. **Arquitectura Monolítica (PHP):** [Monolotico/README.md](Monolotico/README.md)
2. **Microservicio v1 (Go + Docker):** [Microservicio/](Microservicio/)
3. **Microservicio v2 (Go + Kubernetes):** [V2_Microservicio/KUBERNETES_FIXES.md](V2_Microservicio/KUBERNETES_FIXES.md)
4. **Comparación:** [ARQUITECTURAS_COMPARACION.md](ARQUITECTURAS_COMPARACION.md)

### 🚀 Si quieres Ejecutar las Aplicaciones
- **Monolítica:** `cd Monolotico/ && docker-compose up -d`
- **Microservicio v1:** `cd Microservicio/ && docker-compose up -d`
- **Microservicio v2:** `cd V2_Microservicio/ && kubectl apply -f k8s-*.yaml`

---

## 📁 Estructura del Repositorio

```
Go_Lenguaje/
│
├── 📖 README.md ⭐ LEER PRIMERO
│   Descripción general del proyecto y su evolución
│
├── 📚 DOCUMENTACIÓN DE REFERENCIA
│   ├── ARQUITECTURAS_COMPARACION.md ⭐ IMPRESCINDIBLE
│   │   Comparación detallada de las 3 versiones
│   │
│   ├── EJECUCION_COMPLETA.md ⭐ IMPRESCINDIBLE
│   │   Guía para ejecutar todas las versiones
│   │
│   └── IMPLEMENTACION_MONOLITICA.md
│       Resumen de la arquitectura monolítica
│
├── 📁 Practicas/
│   ├── Introduccion.md - Guía de aprendizaje de Go
│   ├── hello-world.go
│   ├── variables.go
│   ├── constantes.go
│   ├── arrays.go
│   ├── slices.go
│   ├── maps.go
│   ├── for.go
│   ├── if.go
│   ├── switch.go
│   └── values.go
│   └─ Propósito: Aprender Go desde cero
│
├── 📁 Monolotico/ ⭐ NUEVA CARPETA
│   ├── 📖 README.md - Documentación completa
│   ├── 🧪 TESTING.md - Casos de prueba
│   ├── ✅ COMPLETADO.md - Resumen de implementación
│   ├── 🚀 quick-start.sh - Script automático
│   │
│   ├── PHP Code (5 archivos)
│   │   ├── index.php
│   │   ├── listar.php
│   │   ├── agregar.php
│   │   ├── editar.php
│   │   └── borrar.php
│   │
│   ├── Configuration (3 archivos)
│   │   ├── config.php
│   │   ├── php.ini
│   │   └── init.sql
│   │
│   ├── Presentation
│   │   └── styles.css
│   │
│   ├── Docker
│   │   └── docker-compose.yml
│   │
│   └─ Propósito: Arquitectura monolítica con PHP + MariaDB
│
├── 📁 Microservicio/
│   ├── main.go - Aplicación monolítica en Go
│   ├── docker-compose.yml
│   ├── Dockerfile
│   ├── db.go
│   ├── altas.go, cambios.go, listas.go
│   ├── index.html, styles.css
│   ├── init.sql
│   └─ Propósito: Versión monolítica en Go con Docker
│
├── 📁 V2_Microservicio/ ⭐ VERSIÓN PRODUCCIÓN
│   ├── 📖 KUBERNETES_FIXES.md - Problemas resueltos
│   │
│   ├── backend/
│   │   ├── main_api.go - API REST
│   │   ├── db.go
│   │   ├── Dockerfile
│   │   └─ Puerto: 8081
│   │
│   ├── frontend/
│   │   ├── main_web.go - Frontend Web
│   │   ├── index.html
│   │   ├── styles.css
│   │   ├── Dockerfile
│   │   └─ Puerto: 8080
│   │
│   ├── kubernetes/ (k8s-*.yaml)
│   │   ├── k8s-db-deployment.yaml
│   │   ├── k8s-api-deployment.yaml
│   │   ├── k8s-web-deployment.yaml
│   │   ├── k8s-secrets.yaml
│   │   └─ Más archivos...
│   │
│   ├── docker-compose.yml
│   ├── init.sql
│   └─ Propósito: Microservicios en Kubernetes
│
├── 📄 Instalacion.md
│   Guía de instalación de herramientas
│
├── 📄 Introduccion.md
│   Introducción general
│
├── 📁 imagenes/
│   Imágenes del proyecto
│
└── 📁 main/
    Contenido adicional
```

---

## 🎯 Guías por Caso de Uso

### Caso 1: Aprender Go Básico
```
1. Leer: Instalacion.md
2. Leer: Practicas/Introduccion.md
3. Ejecutar: Ejercicios en Practicas/
4. Experimento: Modificar ejemplos
```

### Caso 2: Entender Arquitecturas
```
1. Leer: README.md
2. Leer: ARQUITECTURAS_COMPARACION.md
3. Ejecutar: Cada versión localmente
4. Comparar: Diferencias en el navegador
```

### Caso 3: Ejecutar Aplicación Monolítica
```
1. Leer: Monolotico/README.md
2. Ejecutar: docker-compose up -d
3. Acceder: http://localhost:8080
4. Probar: Casos en Monolotico/TESTING.md
```

### Caso 4: Ejecutar Microservicios v1
```
1. Leer: Microservicio/README.md (si existe)
2. Ejecutar: docker-compose up -d
3. Acceder: http://localhost:8080
4. API: http://localhost:8081/api/usuarios
```

### Caso 5: Ejecutar Microservicios v2 (Kubernetes)
```
1. Leer: V2_Microservicio/KUBERNETES_FIXES.md
2. Requisitos: minikube, kubectl
3. Ejecutar: kubectl apply -f k8s-*.yaml
4. Port-forward: kubectl port-forward svc/web-service 8080:8080
5. Acceder: http://localhost:8080
```

### Caso 6: Ejecutar Todas Simultáneamente
```
1. Leer: EJECUCION_COMPLETA.md
2. Usar 3 terminales separadas
3. Terminal 1: Monolotico (puerto 8080)
4. Terminal 2: Microservicio (puerto 8081-8082)
5. Terminal 3: Kubernetes (con port-forward)
6. Comparar: Comportamiento de cada una
```

---

## 📊 Resumen de Versiones

| Versión | Ubicación | Tecnología | Estado |
|---------|-----------|-----------|--------|
| **Monolítica** | `Monolotico/` | PHP 8.2 + MariaDB | ✅ Nueva |
| **Microservicio v1** | `Microservicio/` | Go + Docker Compose | ✅ Existente |
| **Microservicio v2** | `V2_Microservicio/` | Go + Kubernetes | ✅ Producción |

---

## 🔗 Enlaces Rápidos

### Documentos Principales
- [README.md](README.md) - Visión general del proyecto
- [ARQUITECTURAS_COMPARACION.md](ARQUITECTURAS_COMPARACION.md) - Comparativa de arquitecturas
- [EJECUCION_COMPLETA.md](EJECUCION_COMPLETA.md) - Guía de ejecución múltiple

### Monolítica (PHP)
- [Monolotico/README.md](Monolotico/README.md) - Documentación completa
- [Monolotico/TESTING.md](Monolotico/TESTING.md) - Casos de prueba
- [Monolotico/COMPLETADO.md](Monolotico/COMPLETADO.md) - Resumen

### Microservicios
- [V2_Microservicio/KUBERNETES_FIXES.md](V2_Microservicio/KUBERNETES_FIXES.md) - Problemas y soluciones

### Aprendizaje
- [Instalacion.md](Instalacion.md) - Instalación de herramientas
- [Practicas/Introduccion.md](Practicas/Introduccion.md) - Introducción a Go

---

## 🚀 Quick Start Commands

```bash
# Monolítica
cd Monolotico && docker-compose up -d

# Microservicio v1
cd Microservicio && docker-compose up -d

# Microservicio v2
cd V2_Microservicio && minikube start && kubectl apply -f k8s-*.yaml

# Ver todas las versiones
# Ver EJECUCION_COMPLETA.md para ejecutarlas simultáneamente
```

---

## 📖 Lecturas Recomendadas en Orden

1. **Primero:** [README.md](README.md) - Entender qué es el proyecto
2. **Segundo:** [IMPLEMENTACION_MONOLITICA.md](IMPLEMENTACION_MONOLITICA.md) - Ver qué se creó
3. **Tercero:** [ARQUITECTURAS_COMPARACION.md](ARQUITECTURAS_COMPARACION.md) - Comparar versiones
4. **Cuarto:** [EJECUCION_COMPLETA.md](EJECUCION_COMPLETA.md) - Ejecutar y experimentar
5. **Quinto:** Carpeta específica que te interese

---

## ✅ Verificación de Integridad

- [x] Carpeta Monolotico creada con 13 archivos
- [x] 3 documentos de referencia principales
- [x] README.md actualizado con evolución
- [x] Todos los CRUD implementados
- [x] Docker Compose configurado
- [x] Kubernetes manifests listos
- [x] Documentación exhaustiva
- [x] Casos de prueba documentados
- [x] Scripts automáticos incluidos
- [x] Comparativas detalladas

---

## 🎓 Propósito Educativo

Este repositorio demuestra **la evolución completa** de una arquitectura:

```
Monolítica (PHP)
    ↓ (Crece el proyecto)
Microservicios v1 (Docker)
    ↓ (Requiere HA)
Microservicios v2 (Kubernetes)
    ↓ (Escala distribuida)
Fase 3: Distribuido en RPi + x86
```

---

## 📞 Cómo Usar Este Índice

- **Busca por problema:** Usa Ctrl+F para encontrar keywords
- **Busca por tecnología:** PHP, Go, Docker, Kubernetes
- **Busca por carpeta:** Monolotico, Microservicio, V2_Microservicio
- **Busca por acción:** up, down, logs, deploy

---

## 🎉 Estado Final

✅ **Completo y Operacional**
- Todas las versiones implementadas
- Todas las documentaciones creadas
- Todos los ejercicios listos
- Todos los casos de prueba documentados
- Listo para aprender, experimentar y enseñar

---

**Última actualización:** 27 de Diciembre, 2025  
**Creado por:** Sambo  
**Propósito:** Educación en Arquitectura de Microservicios

---

*Para comenzar: Abre [README.md](README.md) o elige una ruta arriba* 🚀
