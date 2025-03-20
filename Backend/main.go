package main

import (
	"Backend/connection"
	"Backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(connection.ConfigCors())
	connection.InitDB()
	// Nueva ruta para obtener datos
	r.GET("/data", controllers.GetData)

	// Iniciar servidor en el puerto 8082
	r.Run(":8082")
}
