package models

const UserTableName = "users"

type User struct {
	ID        string `db:"id"`
	FullName  string `db:"full_name"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int64  `db:"age"`
}

func (u User) TableName() string {
	return UserTableName
}
