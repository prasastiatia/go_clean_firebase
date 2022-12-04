package student

import (
	"remake_crud_golang/app/repository/student"
)

func ListStudent() (interface{}, error) {
	data, err := student.ListStudent()

	if err != nil {
		return nil, err
	}

	return data, nil
}
