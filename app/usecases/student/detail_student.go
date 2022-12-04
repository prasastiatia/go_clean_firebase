package student

import "remake_crud_golang/app/repository/student"

type (
	ParamsDetail struct {
		ID string
	}
)

func DetailStudent(p *ParamsDetail) (map[string]interface{}, error) {
	data, err := student.FindById(&student.ParamsAdd{
		ID: p.ID,
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
