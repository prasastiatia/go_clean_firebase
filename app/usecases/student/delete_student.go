package student

import "remake_crud_golang/app/repository/student"

func DeleteStudent(p *ParamsDetail) error {
	err := student.DeleteById(&student.ParamsAdd{
		ID: p.ID,
	})

	if err != nil {
		return err
	}

	return nil
}
