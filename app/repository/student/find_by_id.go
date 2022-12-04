package student

import (
	"context"
	"fmt"
	firestore "remake_crud_golang/app/drivers"
)

func FindById(p *ParamsAdd) (map[string]interface{}, error) {
	ctx := context.Background()
	client, err := firestore.Firebase(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(p.ID)
	data, err := client.Collection("students").Doc(p.ID).Get(ctx)
	if err != nil {
		return nil, err
	}

	defer client.Close()

	fmt.Println(data.Data())

	return data.Data(), nil
}
