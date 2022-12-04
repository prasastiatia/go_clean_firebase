package student

import (
	"remake_crud_golang/app/entities"
	"remake_crud_golang/app/repository/student"

	"github.com/google/uuid"
)

//Data Parameter tambah data catalog
type Address struct {
	State string
	City  string
}

type (
	ParamsAdd struct {
		Name    string
		Age     int64
		Address Address
	}
)

//Logic untuk menambah data Student
func CreateStudent(p *ParamsAdd) (*entities.Student, error) {
	id := uuid.New().String()
	data, err := student.Create(&student.ParamsAdd{
		ID:   id,
		Name: p.Name,
		Age:  p.Age,
		Address: student.Address{
			State: p.Address.State,
			City:  p.Address.City,
		},
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
