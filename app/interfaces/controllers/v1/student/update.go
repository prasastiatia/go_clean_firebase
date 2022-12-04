package studentController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	usecaseStudent "remake_crud_golang/app/usecases/student"
)

func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		validate := new(createRequest)
		if err := c.Bind(validate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
		}
		err := usecaseStudent.UpdateStudent(&usecaseStudent.ParamsUpdate{
			ID:   id,
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
			"message": "Update Student Berhasil",
		})
	}
}
