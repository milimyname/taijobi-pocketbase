package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// Fetch all feedbacks with their IDs and userIds
		type Feedback struct {
			ID     string `db:"id"`
			UserID string `db:"userId"`
		}
		var feedbacks []Feedback
		err := db.Select("id", "userId").From("feedbacks").All(&feedbacks)
		if err != nil {
			return fmt.Errorf("error fetching feedbacks: %v", err)
		}

		// Group feedbacks by userId
		userFeedbacks := make(map[string][]string)
		for _, feedback := range feedbacks {
			userFeedbacks[feedback.UserID] = append(userFeedbacks[feedback.UserID], feedback.ID)
		}

		// For each userId, create a user if it doesn't exist and update with feedback IDs
		for userId, feedbackIds := range userFeedbacks {
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

			// Update user with feedback IDs
			feedbackIdsJSON, err := json.Marshal(feedbackIds)
			if err != nil {
				return fmt.Errorf("error marshaling feedback IDs: %v", err)
			}
			_, err = db.NewQuery("UPDATE users SET feedbacks = {:feedbacks} WHERE id = {:userId}").
				Bind(dbx.Params{
					"userId":    userId,
					"feedbacks": string(feedbackIdsJSON),
				}).
				Execute()
			if err != nil {
				return fmt.Errorf("error updating user feedbacks: %v", err)
			}
			fmt.Printf("Updated user %s with %d feedbacks\n", userId, len(feedbackIds))
		}

		fmt.Println("Migration completed successfully.")
		return nil
	}, func(db dbx.Builder) error {
		// Down migration: Remove feedbacks field from users and delete users without email
		_, err := db.NewQuery("UPDATE users SET feedbacks = NULL").Execute()
		if err != nil {
			return fmt.Errorf("error removing feedbacks from users: %v", err)
		}

		_, err = db.NewQuery("DELETE FROM users WHERE email = ''").Execute()
		if err != nil {
			return fmt.Errorf("error removing users without email: %v", err)
		}

		fmt.Println("Down migration completed successfully.")
		return nil
	})
}