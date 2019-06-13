package models

type Feedback struct {
	tableName        struct{} `sql:"resident.feedback,alias:feedback"`
	Id               string   `json:"Id" sql:"id"`
	ResidentId       string   `json:"ResidentId" sql:"resident_id"`
	RoomId           string   `json:"RoomId" sql:"room_id"`
	Title            string   `json:"Title" sql:"title"`
	Content          string   `json:"Content" sql:"content"`
	Images           string   `json:"Images" sql:"images"`
	Status           int      `json:"Status" sql:"status"`
	AssignEmployeeId string   `json:"AssignEmployeeId" sql:"employee_id"`
	DueDate          int      `json:"DueDate" sql:"due_date"`
	ActualFinishDate int      `json:"ActualFinishDate" sql:"actual_finish_date"`
	SideId           string   `json:"SideId" sql:"side_id"`
	BlockId          string   `json:"BlockId" sql:"block_id"`
	PositionNote     string   `json:"PositionNote" sql:"position_note"`
	AssignedBy       string   `json:"AssignedBy" sql:"assigned_by"`
}

type FeedbackModel struct {
	tableName        struct{} `sql:"resident.feedback,alias:feedback"`
	Id               string   `json:"Id" sql:"id"`
	ResidentId       string   `json:"ResidentId" sql:"resident_id"`
	RoomId           string   `json:"RoomId" sql:"room_id"`
	Room             Room
	Title            string `json:"Title" sql:"title"`
	Content          string `json:"Content" sql:"content"`
	Images           string `json:"Images" sql:"images"`
	Status           int    `json:"Status" sql:"status"`
	AssignEmployeeId string `json:"AssignEmployeeId" sql:"employee_id"`
	Employee         Employee
	DueDate          int    `json:"DueDate" sql:"due_date"`
	ActualFinishDate int    `json:"ActualFinishDate" sql:"actual_finish_date"`
	SideId           string `json:"SideId" sql:"side_id"`
	Side             Side
	BlockId          string `json:"BlockId" sql:"block_id"`
	Block            Block
	PositionNote     string `json:"PositionNote" sql:"position_note"`
	AssignedBy       string `json:"AssignedBy" sql:"assigned_by"`
}