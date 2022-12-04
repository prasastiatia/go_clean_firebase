package student

import "remake_crud_golang/app/repository/student"

type (
	ParamsUpdate struct {
		ID      string
		Name    string
		Age     int64
		Address Address
	}
)

func UpdateStudent(p *ParamsUpdate) error {

	err := student.UpdateById(&student.ParamsUpdate{
		ID:   p.ID,
		Name: p.Name,
		Age:  p.Age,
		Address: student.Address{
			State: p.Address.State,
			City:  p.Address.City,
		}})

	if err != nil {
		return err
	}

	return nil
}
