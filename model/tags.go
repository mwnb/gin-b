package model

type Tags struct {
	Tag string `json:"tag"`
}

func (*Tags) TableName() string {
	return "tags"
}
