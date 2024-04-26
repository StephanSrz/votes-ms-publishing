# Guía de Descarga y Configuración del Microservicio

## Pasos

1. Clonar y abrir el repositorio.
2. Descargar las dependencias con el comando `go mod tidy`.
3. En este ejemplo y los siguientes se usará un cluster de MongoDB Atlas para usarlo como persistencia de datos. Para poder hacerlo, inicie sesión o cree una cuenta en [MongoDB Atlas](https://account.mongodb.com/account/login?n=https%3A%2F%2Fcloud.mongodb.com%2Fv2&nextHash=%23clusters&signedOut=true).
4. Para obtener la cadena de conexión, haga clic en "Database" y después en "Connect".
5. Siga los pasos para conectarse mediante MongoDB Compass.
6. Para configurar las variables de entorno, hay un archivo llamado `.env.example`, el cual sirve como ejemplo para poder configurar las variables de entorno, usando la cadena de conexión obtenida en el paso anterior.
7. Cree un archivo `.env` y ponga los mismos nombres de las variables con los datos de su cluster.
8. Inicie el servidor con el comando: `go run cmd/api/main.go`.
9. Para poder realizar solicitudes hacia el servidor se usará Postman, el servidor describe las rutas disponibles y los métodos de solicitud. Ya que está corriendo en local la ruta en el puerto 8080 la ruta es: `localhost:8080/...`.
Ruta: localhost:8080/v1/votes
Cuerpo
{
    "userId": "123",
    "roomId": "456",
    "userStory": "US 1",
    "value": "2"
}
