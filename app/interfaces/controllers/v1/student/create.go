package studentController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	usecaseStudent "remake_crud_golang/app/usecases/student"
)

// Fungsi untuk meneruskan data ke users dengan sesuai format
func CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		validate := new(createRequest)
		if err := c.Bind(validate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
		}

		data, err := usecaseStudent.CreateStudent(&usecaseStudent.ParamsAdd{
			Name: validate.Name,
			Age:  validate.Age,
			Address: usecaseStudent.Address{
				State: validate.Address.State,
				City:  validate.Address.City,
			},
		})

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

type Address struct {
	State string `json:"state"`
	City  string `json:"city"`
}

type (
	createRequest struct {
		Name    string  `json:"name"`
		Age     int64   `json:"age"`
		Address Address `json:"address"`
	}
)
