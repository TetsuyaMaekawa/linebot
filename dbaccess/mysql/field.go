package mysql

// Mytable DBのフィールドを保持
type Mytable struct {
	ID   int
	Name string
}

// TableName ...
func (mytable *Mytable) TableName() string {
	return "mytable"
}
