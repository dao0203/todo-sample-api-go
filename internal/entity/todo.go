package entity

type Todo struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CategoryId  int     `json:"category_id"`
	IsCompleted bool    `json:"is_completed"`
	DueDate     int     `json:"due_date"`
}
