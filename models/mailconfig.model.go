package models

type MailConfig struct {
	tableName  struct{} `sql:"resident.mailconfig,alias:mailconfig"`
	Id         int      `json:"Id" sql:"id"`
	Server     string   `json:"Server" sql:"server"`
	UserAuth   int      `json:"UseAuth" sql:"use_auth"`
	UserSecure int      `json:"UseSecure" sql:"use_secure"`
	Username   string   `json:"Username" sql:"username"`
	Password   string   `json:"Password" sql:"password"`
	Port       int      `json:"Port" sql:"port"`
	MailType   string   `json:"MailType" sql:"mail_type"`
}
