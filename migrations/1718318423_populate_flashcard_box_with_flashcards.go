package migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		_, err := db.NewQuery(`UPDATE flashcardBoxes
								SET flashcards = COALESCE(
									(SELECT '[' || GROUP_CONCAT('"' || f.id || '"', ',') || ']'
									FROM flashcard f
									WHERE f.flashcardBox = flashcardBoxes.id),
									'[]'
								);`).Execute()
		return err
	}, nil)
}
