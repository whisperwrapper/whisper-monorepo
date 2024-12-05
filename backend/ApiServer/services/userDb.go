package services

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO make those lines consistent with implementation
type UserDbModel interface {
	getUserInfo(string) (string, string, error)
	isUserInDatabase(string) (bool, error)
	addUserToDatabase(string, string) error
	getUserTranscriptionHistory(string) (pgx.Rows, error)

	saveTranscription(text string, username string, is_translation bool, language string)
	editTranscription(id int, username string, new_content string) error
	deleteTranscription(id int, username string) error

	saveDiarization(text []string, speaker []string, username string, language string)

	saveTranslation(text string, username string, is_translation bool, language string, translated_text string, translation_language string)
}

type UserDb struct {
	pool *pgxpool.Pool
}

func NewUserDb(pool *pgxpool.Pool) *UserDb {
	p := &UserDb{}
	p.pool = pool
	return p
}

func (db UserDb) saveTranscription(text string, user_id string, is_translation bool, language string) {
	_, err := db.pool.Exec(context.Background(), `
    INSERT INTO transcription(app_user_id, content, is_translation, lang) 
    VALUES $1, $2, $3, $4);
	`, user_id, text, is_translation, language)
	if err != nil {
		log.Fatal(err)
	}
}

func (db UserDb) editTranscription(id int, user_id string, new_content string) error {
	_, err := db.pool.Exec(context.Background(), `
    UPDATE transcription set content = $1 where id = $2 and app_user_id = $3;
	`, new_content, id, user_id)
	return err
}

func (db UserDb) deleteTranscription(id int, user_id string) error {
	_, err := db.pool.Exec(context.Background(), `
    DELETE FROM transcription WHERE id = $1 and app_user_id = $2;
	`, id, user_id)
	return err
}

func (db UserDb) saveTranslation(text string, user_id string, is_translation bool, language string, translated_text string, translation_language string) error {
	transcription_id := 0
	err := db.pool.QueryRow(context.Background(), `
    INSERT INTO transcription(app_user_id, content, is_translation, lang) 
    VALUES ($1, $2, $3, $4)
	RETURNING id;
	`, user_id, text, is_translation, language).Scan(&transcription_id)
	if err != nil {
		return (err)
	}
	_, err = db.pool.Exec(context.Background(), `
    INSERT INTO translation(transcription_id, lang, content) 
    VALUES ($1, $2, $3);
	`, transcription_id, translation_language, translated_text)
	if err != nil {
		return (err)
	}
	return nil
}

func (db UserDb) editTranslation(edit_transcription bool, edit_translation bool, transcription_id string, new_transcription string, new_translation string, user_id string) error {
	if edit_transcription {
		_, err := db.pool.Exec(context.Background(), `
		UPDATE transcription set content = $1 where id = $2 and app_user_id = $3;
		`, new_transcription, transcription_id, user_id)
		return err
	}
	if edit_translation {
		_, err := db.pool.Exec(context.Background(), `
		UPDATE translation set content = $1 where transcription_id = $2 and app_user_id = $3;
		`, new_translation, transcription_id, user_id)
		return err
	}
	return nil
}

func (db UserDb) saveDiarization(text []string, speaker []string, user_id string, language string) {
	diarization_id := 0
	err := db.pool.QueryRow(context.Background(), `
    INSERT INTO diarization(app_user_id, lang) 
    VALUES ($1, $2, $3, $4)
	RETURNING id;
	`, user_id, language).Scan(&diarization_id)
	if err != nil {
		log.Fatal(err)
	}
	for i := range text {
		_, err = db.pool.Exec(context.Background(),
			`INSERT INTO speaker_line(id, diarization_id, speaker, content) 
    		VALUES ($1, $2, $3, $4);`,
			i, diarization_id, speaker[i], text[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func editDiarization(new_content []string, diarization_id string) {

}

func (db UserDb) getUserInfo(username string) (string, string, error) {
	row := db.pool.QueryRow(context.Background(), "SELECT id, password_hash FROM app_user WHERE username=$1;", username)
	var password string
	var database_id string
	err := row.Scan(&database_id, &password)
	return database_id, password, err
}

func (db UserDb) isUserInDatabase(username string) (bool, error) {
	var count int
	err := db.pool.QueryRow(context.Background(), "SELECT COUNT(username) FROM app_user WHERE username=$1;", username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (db UserDb) addUserToDatabase(username string, password string) error {
	password_hash, err := HashPassword(password)
	if err != nil {
		log.Printf("Database error %v", err)
		return err
	}
	_, err = db.pool.Exec(context.Background(), "INSERT INTO app_user(username, password_hash) VALUES ($1, $2);", username, password_hash)
	if err != nil {
		log.Printf("Database error %v", err)
		return err
	}
	return nil
}

func (db UserDb) getUserTranscriptionHistory(user_id string) (pgx.Rows, error) {
	rows, err := db.pool.Query(context.Background(), "SELECT id, content, created_at FROM transcription WHERE app_user_id=$1;", user_id)
	if err != nil {
		log.Printf("Database error %v", err)
		return nil, err
	}
	return rows, nil
}
