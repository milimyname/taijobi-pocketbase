package migrations

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		type FlashcardBox struct {
			ID         string `db:"id"`
			UserID     string `db:"userId"`
			Flashcards string `db:"flashcards"` // String representation of flashcards
		}

		type Search struct {
			ID        string `db:"id"`
			User      string `db:"user"`
			Flashcard string `db:"flashcard"`
		}

		var flashcardBoxes []FlashcardBox

		// Fetch flashcard boxes with their user IDs and flashcards
		err := db.Select("id", "userId", "flashcards").From("flashcardBoxes").All(&flashcardBoxes)
		if err != nil {
			return err
		}

		for _, box := range flashcardBoxes {
			// Assuming the flashcards are stored as a JSON array string
			var flashcardIDs []string
			err := json.Unmarshal([]byte(box.Flashcards), &flashcardIDs)
			if err != nil {
				return fmt.Errorf("error unmarshaling flashcards for box %s: %v", box.ID, err)
			}

			if len(flashcardIDs) > 0 {
				// Create the IN clause with actual values
				inClause := "'" + strings.Join(flashcardIDs, "','") + "'"

				// Construct the full query string
				query := fmt.Sprintf("UPDATE flashcard SET user = '%s' WHERE id IN (%s)",
					strings.ReplaceAll(box.UserID, "'", "''"), // Escape single quotes in UserID
					inClause)

				// Execute the query
				_, err := db.NewQuery(query).Execute()
				if err != nil {
					return fmt.Errorf("error updating flashcards for box %s: %v", box.ID, err)
				}
			}
		}

		var searches []Search

		// Fetch searches with their user IDs
		err = db.Select("id", "user", "flashcard").From("searches").All(&searches)
		if err != nil {
			return err
		}

		for _, search := range searches {
			// Construct the full query string, updating only when userId is empty
			query := fmt.Sprintf("UPDATE flashcard SET user = '%s' WHERE id = '%s' AND user = ''", search.User, search.Flashcard)

			// Execute the query
			_, err := db.NewQuery(query).Execute()
			if err != nil {
				return fmt.Errorf("error updating flashcard for search %s: %v", search.ID, err)
			}
		}

		fmt.Println("Migration completed successfully.")

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...

		return nil
	})
}
