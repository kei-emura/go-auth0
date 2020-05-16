package app

import (
	"encoding/gob"
	"log"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	// Store is to store use session
	Store *sessions.FilesystemStore
)

// Init is to initialize user session
func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		return err
	}

	Store = sessions.NewFilesystemStore("", []byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	return nil
}
