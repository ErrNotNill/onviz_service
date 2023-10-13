package login

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"onviz/DB"
)

func GetClientIDAndSecret() {
	rows, err := DB.Db.Query("SELECT client_id, client_secret, redirect_uri FROM clients")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Process the results
	for rows.Next() {
		var clientID, clientSecret, redirectURI string
		if err := rows.Scan(&clientID, &clientSecret, &redirectURI); err != nil {
			log.Fatal(err)
		}
		log.Printf("Client ID: %s, Client Secret: %s, Redirect URI: %s", clientID, clientSecret, redirectURI)
	}
}

func AssumeHash() {
	// Assuming you've retrieved the user's stored hashed password from the database
	storedHashedPassword := []byte("...") // Retrieve this from the database
	// User's provided password
	userPassword := []byte("user-provided-password")
	// Compare the hashed password with the user's provided password
	err := bcrypt.CompareHashAndPassword(storedHashedPassword, userPassword)
	if err != nil {
		// Passwords do not match, user authentication failed
	} else {
		// Passwords match, user is authenticated
	}
}
