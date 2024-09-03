package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// Fetch all quizzes with their IDs and userIds
		type Quiz struct {
			ID     string `db:"id"`
			UserID string `db:"userId"`  // Assuming the field in quizzes is named "user"
		}
		var quizzes []Quiz
		err := db.Select("id", "userId").From("quizzes").All(&quizzes)
		if err != nil {
			return fmt.Errorf("error fetching quizzes: %v", err)
		}

		// Group quizzes by userId
		userQuizzes := make(map[string][]string)
		for _, quiz := range quizzes {
			userQuizzes[quiz.UserID] = append(userQuizzes[quiz.UserID], quiz.ID)
		}

		// Update users with their quiz IDs
		for userId, quizIds := range userQuizzes {
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

			// Update user with quiz IDs
			quizIdsJSON, err := json.Marshal(quizIds)
			if err != nil {
				return fmt.Errorf("error marshaling quiz IDs: %v", err)
			}
			_, err = db.NewQuery("UPDATE users SET quizzes = {:quizzes} WHERE id = {:userId}").
				Bind(dbx.Params{
					"userId":  userId,
					"quizzes": string(quizIdsJSON),
				}).
				Execute()
			if err != nil {
				return fmt.Errorf("error updating user quizzes: %v", err)
			}
			fmt.Printf("Updated user %s with %d quizzes\n", userId, len(quizIds))
		}

		fmt.Println("Migration completed successfully.")
		return nil
	}, func(db dbx.Builder) error {
		// Down migration: Remove quizzes field from users
		_, err := db.NewQuery("UPDATE users SET quizzes = NULL").Execute()
		if err != nil {
			return fmt.Errorf("error removing quizzes from users: %v", err)
		}

		fmt.Println("Down migration completed successfully.")
		return nil
	})
}