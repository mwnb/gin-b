package model

type Picturn struct {
	Id         int64  `json:"id"`
	Alt        string `json:"alt"`         //描述
	Path       string `json:"path"`        //路径
	Tag        string `json:"tag"`         //标签
	Origin     string `json:"origin"`      //来源
	Size       string `json:"size"`        // 大小
	PageView   string `json:"page_view"`   //浏览数
	Anthor     string `json:"anthor"`      //作者
	UploadTime string `json:"upload_time"` //上传时间
	DeleteAt   string `json:"delete_at"`   // 是否删除
	Star       int64  `json:"start"`       // 点赞
	Trample    int64  `json:"trample"`     // 踩
	Collect    int64  `json:"collect"`     // 收藏数量
	Related    int64  `json:"related"`     //相关图片数量
}

func (Picturn) TableName() string {
	return "picturn"
}
