package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"path/filepath"
)

var INSTANCE *auth.Client

func Init() {
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}
	//Firebase Auth
	authClient, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error")
	}
	INSTANCE = authClient
}
