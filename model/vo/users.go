package vo

type UserInput struct {
	SystemId   int  `json:"system_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Type       int  `json:"type"`
	UpdateTime int
}
