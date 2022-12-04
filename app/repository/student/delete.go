package student

import (
	"context"
	firestore "remake_crud_golang/app/drivers"
)

func DeleteById(p *ParamsAdd) error {
	ctx := context.Background()
	client, err := firestore.Firebase(ctx)
	if err != nil {
		return err
	}

	_, err = client.Collection("students").Doc(p.ID).Delete(ctx)
	if err != nil {
		return err
	}

	defer client.Close()

	return nil
}
