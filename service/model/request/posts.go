package request

type GetPosts struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type InsertPosts struct {
	Title    string `json:"title" valid:"required,min=20" validname:"title"`
	Content  string `json:"content" valid:"required,min=200" validname:"content"`
	Category string `json:"category" valid:"required,min=3" validname:"category"`
	Status   string `json:"status" valid:"required,oneof=Publish|Draft|Thrash" validname:"status"`
}

type UpdatePosts struct {
	Id       string `json:"id" valid:"required" validname:"id"`
	Title    string `json:"title" valid:"required,min=20" validname:"title"`
	Content  string `json:"content" valid:"required,min=200" validname:"content"`
	Category string `json:"category" valid:"required,min=3" validname:"category"`
	Status   string `json:"status" valid:"required,oneof=Publish|Draft|Thrash" validname:"status"`
}
