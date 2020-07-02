package vo

type UserInput struct {
	SystemId   int32  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Type       int32  `json:"type"`
	UpdateTime int64
}
