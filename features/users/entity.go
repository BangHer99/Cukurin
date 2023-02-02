package users

type Core struct {
	ID       uint   `json:"ID"`
	Name     string `json:"name"`
	Phone    int    `json:"phone"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

type Cukuran struct {
	UserID      uint
	Title       string
	Image       string
	Description string
	Price       int
}

type UsecaseInterface interface {
	Create(data Core) (row int, err error)
}

type DataInterface interface {
	Register(data Core) (row int, err error)
}
