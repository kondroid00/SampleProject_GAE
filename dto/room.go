package dto

type Room struct {
	Id        string `datastore:"-" goon:"id" json:"id"`
	Name      string `json:"name"`
	Theme     string `json:"theme"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
