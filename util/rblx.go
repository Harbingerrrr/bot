package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// Get group ID from server ID
func GetGroupID(serverID *string) int {
	// get group ID from database
	return 2885688
}

// Get user IDs from usernames
func GetUserIDsFromUsernames(usernames *[]string) []int {
	var userIDs []int

	client := &http.Client{}
	for _, user := range *usernames {

		payload, err := getUserIDsFromUsernames_Payload(user)
		if err != nil {
			fmt.Printf("Error creating payload for username %s: %v\n", user, err)
			continue
		}

		body, err := GetUserIDsFromUsernames_Request(client, payload)
		if err != nil {
			log.Printf("Error making request for username %s: %v", user, err)
			continue
		}

		userIDs = append(userIDs, int(gjson.Get(string(body), "data.0.id").Int()))
	}

	return userIDs
}

// Create payload for usernames
func getUserIDsFromUsernames_Payload(userName string) ([]byte, error) {
	payload := map[string]interface{}{
		"usernames":          []string{userName},
		"excludeBannedUsers": true,
	}
	return json.Marshal(payload)
}

// Make request to get user IDs from usernames
func GetUserIDsFromUsernames_Request(client *http.Client, payload []byte) ([]byte, error) {
	resp, err := client.Post("https://users.roblox.com/v1/usernames/users", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("POST request failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}

func QuickCheck(userID *int, groupID *int) {
	// check if user can promote
	// if user can promote, promote user
	// if user can't promote, return
}
