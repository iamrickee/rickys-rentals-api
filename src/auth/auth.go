package auth

type Resp struct {
	Profile User   `json:"profile"`
	Message string `json:"message"`
}

type User struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
