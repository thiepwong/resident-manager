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

type NotificationModel struct {
	tableName   struct{} `sql:"resident.send_notification,alias:send_notification"`
	Id          string   `json:"Id" sql:"id"`
	SideId      string   `json:"SideId" sql:"side_id"`
	Side        Side
	Title       string `json:"Title" sql:"title"`
	PublishDate int64  `json:"PublishDate" sql:"publish_date"`
	SendResult  bool   `json:"SendResult" sql:"send_success"`
	Content     string `json:"Content" sql:"body"`
}

type NotificationQueryModel struct {
	Id          string `json:"Id" sql:"id"`
	Title       string `json:"Title" sql:"title"`
	PublishDate int64  `json:"PublishDate" sql:"publish_date"`
	SendResult  bool   `json:"SendResult" sql:"send_success"`
	Content     string `json:"Content" sql:"body"`
	SideId      string `json:"SideId" sql:"side_id"`
	SideName    string `json:"SideName" sql:"side_name"`
	Hotline     string `json:"Hotline" sql:"hotline"`
	Address     string `json:"Address" sql:"address"`
	CoverPhotos string `json:"CoverPhotos" sql:"cover_photos"`
}

type SendNotification struct {
	Topic   string `json:"Topic"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}
