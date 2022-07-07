package model

type Images struct {
	Id         int64  `json:"id"`
	ImgPath    string `json:"img_path"`
	Tags       string `json:"tags"`
	CreateTime string `json:"create_time"`
}

func (*Images) TableName() string {
	return "images"
}
