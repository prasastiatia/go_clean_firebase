package student

import (
	"context"
	"remake_crud_golang/app/entities"

	firestore "remake_crud_golang/app/drivers"
)

//Data Parameter tambah data student

type (
	ParamsAdd struct {
		ID      string
		Name    string
		Age     int64
		Address Address
	}
)

//Query untuk Tambah data ke database Firestore
func Create(p *ParamsAdd) (*entities.Student, error) {
	ctx := context.Background()
	client, err := firestore.Firebase(ctx)
	if err != nil {
		return nil, err
	}

	Address := entities.Address{
		State: p.Address.State,
		City:  p.Address.City,
	}

	Student := entities.Student{
		ID:      p.ID,
		Name:    p.Name,
		Age:     p.Age,
		Address: Address,
	}

	_, err = client.Collection("students").Doc(p.ID).Set(ctx, &Student)
	if err != nil {
		return nil, err
	}

	defer client.Close()

	return entities.ToStudentEntity(&Student), nil
}
