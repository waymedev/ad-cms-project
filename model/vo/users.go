package vo

type UserInput struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     int32  `json:"type"`
}
