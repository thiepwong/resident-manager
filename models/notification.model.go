package models

type Notification struct {
	tableName   struct{} `sql:"resident.send_notification,alias:send_notification"`
	Id          string   `json:"Id" sql:"id"`
	SideId      string   `json:"SideId" sql:"side_id"`
	Title       string   `json:"Title" sql:"title"`
	PublishDate int64    `json:"PublishDate" sql:"publish_date"`
	SendResult  bool     `json:"SendResult" sql:"send_success"`
	Content     string   `json:"Content" sql:"body"`
}
