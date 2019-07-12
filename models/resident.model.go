package models

type Resident struct {
	tableName   struct{} `sql:"resident.resident,alias:resident"`
	Id          string   `json:"Id" sql:"id"`
	AccountId   string   `json:"AccountId" sql:"account_id"`
	PhoneNumber string   `json:"PhoneNumber" sql:"phone_no"`
	FullName    string   `json:"FullName" sql:"full_name"`
	Email       string   `json:"Email" sql:"email"`
}

type ResidentRoom struct {
	tableName  struct{} `sql:"resident.resident_room_mapping,alias:resident_room_mapping"`
	Id         string   `json:"Id" sql:"id"`
	ResidentId string   `json:"ResidentId" sql:"resident_id"`
	Resident   Resident
	RoomId     string `json:"RoomId" sql:"room_id"`
	Room       RoomModel
	Active     bool `json:"Active" sql:"active"`
	Deleted    bool `json:"Deleted" sql:"deleted"`
}

type ResidentModel struct {
	Resident       Resident `json:"Resident"`
	RoomId         string   `json:"RoomId"`
	ResidentRoomId string   `json:"ResidentRoomId"`
}
