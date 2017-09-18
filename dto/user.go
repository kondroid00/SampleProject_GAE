package dto

type User struct {
	Id   string `datastore:"-" goon:"id" json:"id"`
	Name string `json:"name"`
}
