package models

type Auth struct {
	IssueId     string `json:"iss"`
	Issuer      string `json:"isn"`
	IssueAt     int    `json:"iat"`
	Expired     int    `json:"exp"`
	System      string `json:"sys"`
	Username    string `json:"usr"`
	CreatedDate int    `json:"urd"`
	IsEnable    int    `json:"ist"`
}

type Signin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	System   string `json:"system"`
}
