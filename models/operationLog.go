package models

type OperationInfo struct {
	UserID    int64       `json:"user_id"`
	Operation string      `json:"operation"`
	Time      string      `json:"time"`
	Status    string      `json:"status"`
	Data      interface{} `json:"data"`
}
