package entity

// by default gorm convert camelcase (TestTable) become snakecase and lowercase (new_price)
// and name of type struct become plurar (test_tables)

// in v2 gorm we can change behaviour of conversion to communicate with database table

// package for defined mapping database
type TestTable struct {
	ID          int    `json:"id"`
	Name        string `json:"name" faker:"name"`
	Description string `json:"description" faker:"url"`
}
