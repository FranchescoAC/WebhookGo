# Webhook Server en Go

Este proyecto implementa un servidor de Webhook básico en Go con soporte para los métodos HTTP `POST` y `GET`. Es ideal para comprender cómo manejar solicitudes entrantes en un servidor y almacenar datos temporalmente en memoria.

## Características

- **POST**: Recibe datos enviados al webhook y los almacena en memoria.
- **GET**: Devuelve los datos almacenados actualmente en memoria.
- **Servidor HTTP**: Configuración y ejecución de un servidor web básico en Go.

## Requisitos previos

- Go instalado (versión 1.16 o superior).
- Editor de código (Visual Studio Code, GoLand, etc.).
- Conexión a internet para instalar dependencias si es necesario.

## Instalación

1. Clona este repositorio:

   ```bash
   git clone https://github.com/tuusuario/webhook-go.git
   cd webhook-go
2. Corre el programa 
- go run main.go
3. Envia datos al servidor
- Invoke-RestMethod -Uri http://localhost:8080/webhook -Method POST -Body '{"message":"Hello, Webhook!"}' -ContentType "application/json"

4. Respuesta esperada
- {
    "status": "received",
    "message": "Hola desde POST"
}
