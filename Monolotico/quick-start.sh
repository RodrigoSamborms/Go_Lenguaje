#!/bin/bash

# Quick Start Script - Monolítica (PHP)
# Uso: bash quick-start.sh

echo "🚀 Iniciando Arquitectura Monolítica..."
echo ""
echo "Requisitos:"
echo "✓ Docker Desktop instalado"
echo "✓ Docker Compose instalado"
echo ""

# Verificar Docker
if ! command -v docker &> /dev/null; then
    echo "❌ Docker no está instalado"
    exit 1
fi

# Verificar Docker Compose
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose no está instalado"
    exit 1
fi

echo "✅ Verificación de requisitos completada"
echo ""

# Detener contenedores previos si existen
echo "🛑 Deteniendo contenedores previos..."
docker-compose down -v 2>/dev/null || true
echo ""

# Levantar servicios
echo "📦 Levantando servicios..."
docker-compose up -d

echo ""
echo "⏳ Esperando inicialización de base de datos (10 segundos)..."
sleep 10

# Verificar que esté activo
echo ""
echo "✅ Verificando servicios..."
docker-compose ps

echo ""
echo "════════════════════════════════════════"
echo "🎉 ¡Aplicación Monolítica Iniciada!"
echo "════════════════════════════════════════"
echo ""
echo "📍 Acceso a la aplicación:"
echo "   URL: http://localhost:8080"
echo ""
echo "📊 Comandos útiles:"
echo "   Ver logs:     docker-compose logs -f web"
echo "   Detener:      docker-compose down"
echo "   Reiniciar:    docker-compose restart"
echo ""
echo "📚 Documentación:"
echo "   README.md    - Guía completa"
echo "   TESTING.md   - Casos de prueba"
echo ""
