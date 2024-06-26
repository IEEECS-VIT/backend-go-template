package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client
func InitFirebase() {
    ctx := context.Background()
    
    opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CREDENTIALS")))
    
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("Error initializing Firebase app: %v\n", err)
    }
    
    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalf("Error initializing Firestore client: %v\n", err)
    }
    
    FirestoreClient = client
    log.Println("Firebase initialized successfully")
}