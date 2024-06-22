package migrations

import (
	"fmt"
	"strings"

	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func classifyWord(word string) string {
	verbEndings := []string{"う", "く", "す", "つ", "ぬ", "ふ", "む", "る", "ぐ", "ぶ"}
	masuEnding := "ます"
	teEnding := "て"
	taEnding := "た"
	extendedVerbEndings := []string{
		"いる", "える", "れる", "せる", "てる", "ける", "ねる", "べる", "める", "るる",
	}

	iAdjectiveEnding := "い"
	naAdjectiveHint := "な"

	runeWord := []rune(word)
	lastCharacter := string(runeWord[len(runeWord)-1])
	secondLastCharacter := ""
	if len(runeWord) > 1 {
		secondLastCharacter = string(runeWord[len(runeWord)-2])
	}

	isVerb := false
	for _, ending := range verbEndings {
		if lastCharacter == ending {
			isVerb = true
			break
		}
	}
	if !isVerb {
		for _, ending := range extendedVerbEndings {
			if strings.HasSuffix(word, ending) {
				isVerb = true
				break
			}
		}
	}
	if isVerb || strings.HasSuffix(word, masuEnding) || strings.HasSuffix(word, teEnding) || strings.HasSuffix(word, taEnding) {
		return "verb"
	}

	if lastCharacter == iAdjectiveEnding && secondLastCharacter != naAdjectiveHint {
		return "adjective"
	}
	if strings.HasSuffix(word, naAdjectiveHint) || (secondLastCharacter == naAdjectiveHint && lastCharacter == iAdjectiveEnding) {
		return "adjective"
	}

	return "unknown"
}

func init() {
	m.Register(func(db dbx.Builder) error {
		// add up queries...

		var flashcards []struct {
			ID   string `db:"id"`
			Word string `db:"name"`
		}

		err := db.Select("id", "name").From("flashcard").All(&flashcards)
		if err != nil {
			return err
		}

		for _, flashcard := range flashcards {
			partOfSpeech := classifyWord(flashcard.Word)
			_, err := db.Update("flashcard", dbx.Params{"partOfSpeech": partOfSpeech}, dbx.HashExp{"id": flashcard.ID}).Execute()
			if err != nil {
				return err
			}
		}

		fmt.Print("Migration completed successfully.")

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...
		// If needed, you can define logic to revert the changes made in the up query.
		// _, err := db.Update("flashcard", dbx.Params{"partOfSpeech": ""}).Execute()
		// if err != nil {
		// 	return err
		// }

		return nil
	})
}
