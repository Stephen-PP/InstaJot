package models

type NoteTag struct {
	ID        int64  `json:"id"`
	NoteID    int64  `json:"note_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	UserID    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
