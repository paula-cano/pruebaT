package main

import (
	"github.com/gin-gonic/gin"
	"Backend/controllers"
)

func main() {
	r := gin.Default()

	// Nueva ruta para obtener datos
	r.GET("/data", controllers.GetData)

	// Iniciar servidor en el puerto 8082
	r.Run(":8082")
}
