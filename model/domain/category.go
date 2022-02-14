package domain

//setiap tablee di representasikan kedalam struct
type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
