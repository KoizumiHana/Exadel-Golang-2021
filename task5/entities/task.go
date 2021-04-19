package entities

type Status string

const (
	NEW         Status = "NEW"
	IN_PROGRESS Status = "IN_PROGRESS"
	CANCELED    Status = "CANCELED"
	DONE        Status = "DONE"
	EXPIRED     Status = "EXPIRED"
)

type Task struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     string `json:"duedate"`
	Status      Status `json:"status"`
}
