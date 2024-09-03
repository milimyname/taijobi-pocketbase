package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// Fetch all chats with their IDs and userIds
		type Chat struct {
			ID     string `db:"id"`
			UserID string `db:"user"`
		}
		var chats []Chat
		err := db.Select("id", "user").From("chats").All(&chats)
		if err != nil {
			return fmt.Errorf("error fetching chats: %v", err)
		}

		// Group chats by userId
		userChats := make(map[string][]string)
		for _, chat := range chats {
			userChats[chat.UserID] = append(userChats[chat.UserID], chat.ID)
		}

		// For each userId, create a user if it doesn't exist and update with chat IDs
		for userId, chatIds := range userChats {
			var exists int
			err := db.NewQuery("SELECT COUNT(*) FROM users WHERE id = {:userId}").
				Bind(dbx.Params{"userId": userId}).
				Row(&exists)
			if err != nil {
				return fmt.Errorf("error checking user existence: %v", err)
			}

			if exists == 0 {
				// Create new user
				_, err := db.NewQuery("INSERT INTO users (id, created, updated, email, verified) VALUES ({:userId}, NOW(), NOW(), '', FALSE)").
					Bind(dbx.Params{"userId": userId}).
					Execute()
				if err != nil {
					return fmt.Errorf("error inserting new user: %v", err)
				}
				fmt.Printf("Inserted new user with id: %s\n", userId)
			}

			// Update user with chat IDs
			chatIdsJSON, err := json.Marshal(chatIds)
			if err != nil {
				return fmt.Errorf("error marshaling chat IDs: %v", err)
			}
			_, err = db.NewQuery("UPDATE users SET chats = {:chats} WHERE id = {:userId}").
				Bind(dbx.Params{
					"userId": userId,
					"chats":  string(chatIdsJSON),
				}).
				Execute()
			if err != nil {
				return fmt.Errorf("error updating user chats: %v", err)
			}
			fmt.Printf("Updated user %s with %d chats\n", userId, len(chatIds))
		}

		fmt.Println("Chat migration completed successfully.")
		return nil
	}, func(db dbx.Builder) error {
		// Down migration: Remove chats field from users
		_, err := db.NewQuery("UPDATE users SET chats = NULL").Execute()
		if err != nil {
			return fmt.Errorf("error removing chats from users: %v", err)
		}

		fmt.Println("Chat down migration completed successfully.")
		return nil
	})
}
