package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// Fetch all flashcard collections with their IDs and userIds
		type FlashcardCollection struct {
			ID     string `db:"id"`
			UserID string `db:"userId"` // Assuming the field in flashcardCollections is named "user"
		}
		var flashcardCollections []FlashcardCollection
		err := db.Select("id", "userId").From("flashcardCollections").All(&flashcardCollections)
		if err != nil {
			return fmt.Errorf("error fetching flashcardCollections: %v", err)
		}

		// Group flashcardCollections by userId
		userFlashcardCollections := make(map[string][]string)
		for _, collection := range flashcardCollections {
			userFlashcardCollections[collection.UserID] = append(userFlashcardCollections[collection.UserID], collection.ID)
		}

		// Update users with their flashcard collection IDs
		for userId, collectionIds := range userFlashcardCollections {
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

			// Update user with collection IDs
			collectionIdsJSON, err := json.Marshal(collectionIds)
			if err != nil {
				return fmt.Errorf("error marshaling flashcard collection IDs: %v", err)
			}
			_, err = db.NewQuery("UPDATE users SET flashcardCollections = {:flashcardCollections} WHERE id = {:userId}").
				Bind(dbx.Params{
					"userId":               userId,
					"flashcardCollections": string(collectionIdsJSON),
				}).
				Execute()
			if err != nil {
				return fmt.Errorf("error updating user flashcard collections: %v", err)
			}
			fmt.Printf("Updated user %s with %d flashcard collections\n", userId, len(collectionIds))
		}

		fmt.Println("Migration completed successfully.")
		return nil
	}, func(db dbx.Builder) error {
		// Down migration: Remove flashcard collections field from users
		_, err := db.NewQuery("UPDATE users SET flashcardCollections = NULL").Execute()
		if err != nil {
			return fmt.Errorf("error removing flashcardCollections from users: %v", err)
		}

		fmt.Println("Down migration completed successfully.")
		return nil
	})
}
