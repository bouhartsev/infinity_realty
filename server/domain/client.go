package domain

type Client struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Telephone  string `json:"telephone"`
	Email      string `json:"email"`
}
