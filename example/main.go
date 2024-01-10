package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/go-identiface"
)

type user struct {
	id        string `json:"id,omitempty"`
	name      string `json:"name,omitempty"`
	imagePath string `json:"image_name,omitempty"`
}

var (
	assetsDir = "./assets"
	modelsDir = filepath.Join(assetsDir, "models")
	imagesDir = filepath.Join(assetsDir, "images")
)

// Run this code at root dir of project with `go run ./example/main.go`
func main() {
	// Imagine this is data users from database
	users := []user{
		{id: uuid.NewString(), name: "Tony Stark", imagePath: filepath.Join(imagesDir, "tony-stark.jpeg")},
		{id: uuid.NewString(), name: "Tzuyu", imagePath: filepath.Join(imagesDir, "tzuyu.jpg")},
	}

	getUserByID := map[string]user{}
	for _, user := range users {
		getUserByID[user.id] = user
	}

	// Initialize `identiface` [string] is data type of key for each image
	iFace, err := identiface.Init[string](modelsDir)
	if err != nil {
		panic(err)
	}

	// Adding image datasets to `go-identiface`
	for _, user := range users {
		fileBytes, err := os.ReadFile(user.imagePath)
		if err != nil {
			panic(err)
		}

		if err := iFace.AddSingleDatasetFromBytes(user.id, fileBytes); err != nil {
			panic(err)
		}
	}

	// After adding datasets, don't forget to load them so they can be used by `go-identiface`. After this you can identify user by human face of image
	iFace.LoadDatasets()

	// Imagine this is an image file sent from client (ex. Mobile/Frontend/Etc)
	clientImg := "tony-stark3.jpeg"
	clienImgBytes, err := os.ReadFile(filepath.Join(imagesDir, clientImg))
	if err != nil {
		panic(err)
	}

	// After get image file from client then you can identify client by face of the image
	data, err := iFace.ClassifySingleFromBytes(clienImgBytes)
	if err != nil { // Will return error if face is not recognized or image contain multiple face
		panic(err)
	}

	// Imagine this is method to get detail data user from db
	user := getUserByID[data.ID]

	// Now you can identify user by face
	fmt.Printf("user: %#v\n", user)
}
