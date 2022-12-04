package student

import (
	"context"
	"log"
	firestore "remake_crud_golang/app/drivers"
	"remake_crud_golang/app/entities"

	"google.golang.org/api/iterator"
)

func ListStudent() (interface{}, error) {
	ctx := context.Background()
	client, err := firestore.Firebase(ctx)
	if err != nil {
		return nil, err
	}

	var students []entities.Student
	it := client.Collection("students").Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		student := entities.Student{
			ID:   doc.Data()["ID"].(string),
			Name: doc.Data()["Name"].(string),
			Age:  doc.Data()["Age"].(int64),
		}
		students = append(students, student)

	}

	return students, nil
}
