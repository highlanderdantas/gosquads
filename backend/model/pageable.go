package model

//Pageable paginação
type Pageable struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

//Content estrutura paginação
type Content struct {
	Content       interface{} `json:"content"`
	TotalElements int         `json:"totalElements"`
	Pageable      Pageable    `json:"pageable"`
}
