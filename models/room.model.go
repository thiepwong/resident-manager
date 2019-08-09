package models

type Room struct {
	tableName struct{} `sql:"resident.room,alias:room"`
	Id        string   `json:"Id" sql:"id"`
	RoomNo    string   `json:"RoomNo" sql:"room_no"`
	SideId    string   `json:"SideId" sql:"side_id"`
	BlockId   string   `json:"BlockId" sql:"block_id"`
}

type RoomModel struct {
	tableName struct{} `sql:"resident.room,alias:room"`
	Id        string   `json:"Id" sql:"id"`
	RoomNo    string   `json:"RoomNo" sql:"room_no"`
	SideId    string   `json:"SideId" sql:"side_id"`
	Side      Side
	//	Side      Side
	BlockId string `json:"BlockId" sql:"block_id"`
	Block   BlockModel
}

type RoomQueyModel struct {
	Id        string `json: "Id" sql: "id" `
	RoomNo    string `json:"RoomNo" sql:"room_no"`
	BlockId   string `json:"BlockId" sql:"blockid"`
	BlockName string `json:"BlockName" sql:"blockname"`
	SideId    string `json:"SideId" sql:"sideid"`
	SideName  string `json:"SideName" sql:"sidename"`
}

type RoomResidentModel struct {
	tableName struct{} `sql:"resident.room,alias:room"`
	Id        string   `json:"Id" sql:"id"`
	RoomNo    string   `json:"RoomNo" sql:"room_no"`
	SideId    string   `json:"SideId" sql:"side_id"`
	Side      Side
	//	Side      Side
	BlockId  string `json:"BlockId" sql:"block_id"`
	Block    BlockModel
	Resident []ResidentRoomShort
}
