package structure

type Account struct {
	Seq         int    `json:"seq"`
	Id          string `json:"id"`
	Pw          string `json:"pw"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Hp          string `json:"hp"`
	Role        int    `json:"role"`
	State       int    `json:"state"`
	Description string `json:"description"`
}
