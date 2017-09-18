package dto

type Token struct {
	Token     string `datastore:"-" goon:"id" json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}
