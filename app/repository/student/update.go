package student

import (
	"context"
	firestore "remake_crud_golang/app/drivers"

	fs1 "cloud.google.com/go/firestore"
)

type Address struct {
	State string
	City  string
}

type (
	ParamsUpdate struct {
		ID      string
		Name    string
		Age     int64
		Address Address
	}
)

func UpdateById(p *ParamsUpdate) error {
	ctx := context.Background()
	client, err := firestore.Firebase(ctx)
	if err != nil {
		return err
	}

	Address := Address{
		State: p.Address.State,
		City:  p.Address.City,
	}

	_, err = client.Collection("students").Doc(p.ID).Update(ctx, []fs1.Update{
		{Path: "Name", Value: p.Name},
		{Path: "Age", Value: p.Age},
		{Path: "Address", Value: Address},
	})
	if err != nil {
		return err
	}

	defer client.Close()

	return nil
}
