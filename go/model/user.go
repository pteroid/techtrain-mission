package model

type User struct {
	ID    uint   `db:"ID, primarykey, autoincrement" json:"id"`
	Name  string `db:"Name, notnull" json:"name"`
}
