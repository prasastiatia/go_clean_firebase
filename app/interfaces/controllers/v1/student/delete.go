package studentController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	usecaseStudent "remake_crud_golang/app/usecases/student"
)

func DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := usecaseStudent.DeleteStudent(&usecaseStudent.ParamsDetail{
			ID: id,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Data berhasil dihapus",
		})
	}
}
