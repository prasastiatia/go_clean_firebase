package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

//Deklarasi Firestore untuk kebutuhan repository
func Firebase(ctx context.Context) (*firestore.Client, error) {
	opt := option.WithCredentialsFile("./serviceKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Panicln("error initializing app: ", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return client, nil
}
