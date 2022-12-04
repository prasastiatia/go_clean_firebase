package studentController

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	usecaseStudent "remake_crud_golang/app/usecases/student"
)

func ListStudent() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, err := usecaseStudent.ListStudent()

		j, _ := json.MarshalIndent(data, "", "  ")
		log.Println(string(j))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "OK",
			"data":    data,
		})
	}
}
