package request

type Article struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author  string `json:"author" binding:"required"`
}

func (article Article) GetMessage() ValidatorMessage {
	return ValidatorMessage{
		"Title.required":   "文章名称不能为空",
		"Content.required": "文章内容不能为空",
		"Author.required":  "文章作者不能为空",
	}
}
