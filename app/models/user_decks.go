package models

type UserDecks struct {
	UserId  int64  `json:"user_id,string"`
	DecksId uint64 `json:"decks_id,string"`
}
