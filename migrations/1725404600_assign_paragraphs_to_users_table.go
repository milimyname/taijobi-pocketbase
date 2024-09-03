package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// Fetch all paragraphs with their IDs and userIds
		type Paragraph struct {
			ID     string `db:"id"`
			UserID string `db:"user"`  // Assuming the field in paragraphs is named "user"
		}
		var paragraphs []Paragraph
		err := db.Select("id", "user").From("paragraphs").All(&paragraphs)
		if err != nil {
			return fmt.Errorf("error fetching paragraphs: %v", err)
		}

		// Group paragraphs by userId
		userParagraphs := make(map[string][]string)
		for _, paragraph := range paragraphs {
			userParagraphs[paragraph.UserID] = append(userParagraphs[paragraph.UserID], paragraph.ID)
		}

		// Update users with their paragraph IDs
		for userId, paragraphIds := range userParagraphs {
			// Check if user exists
			var exists int
			err := db.NewQuery("SELECT COUNT(*) FROM users WHERE id = {:userId}").
				Bind(dbx.Params{"userId": userId}).
				Row(&exists)
			if err != nil {
				return fmt.Errorf("error checking user existence: %v", err)
			}

			if exists == 0 {
				fmt.Printf("User with id %s not found, skipping\n", userId)
				continue
			}

			// Update user with paragraph IDs
			paragraphIdsJSON, err := json.Marshal(paragraphIds)
			if err != nil {
				return fmt.Errorf("error marshaling paragraph IDs: %v", err)
			}
			_, err = db.NewQuery("UPDATE users SET paragraphs = {:paragraphs} WHERE id = {:userId}").
				Bind(dbx.Params{
					"userId":     userId,
					"paragraphs": string(paragraphIdsJSON),
				}).
				Execute()
			if err != nil {
				return fmt.Errorf("error updating user paragraphs: %v", err)
			}
			fmt.Printf("Updated user %s with %d paragraphs\n", userId, len(paragraphIds))
		}

		fmt.Println("Migration completed successfully.")
		return nil
	}, func(db dbx.Builder) error {
		// Down migration: Remove paragraphs field from users
		_, err := db.NewQuery("UPDATE users SET paragraphs = NULL").Execute()
		if err != nil {
			return fmt.Errorf("error removing paragraphs from users: %v", err)
		}

		fmt.Println("Down migration completed successfully.")
		return nil
	})
}