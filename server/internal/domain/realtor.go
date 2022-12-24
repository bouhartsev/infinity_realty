package domain

type Realtor struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic string  `json:"patronymic"`
	Commission float32 `json:"commission"`
}
