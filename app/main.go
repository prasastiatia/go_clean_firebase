package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	studentRoutes "remake_crud_golang/app/interfaces/routes/v1"
)

func main() {
	godotenv.Load()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	studentRoutes.MountStudent(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
