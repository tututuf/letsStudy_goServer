package interfaces

type UserInfo struct {
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	Username string `json:"username"`
	Power    int8   `json:"power"`
	password string
}
