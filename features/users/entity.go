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
	AllUsers() (user []Core, err error)
	UpdateUser(data Core) (row int, err error)
	Put(input Core) (data Core, err error)
}

type DataInterface interface {
	Register(data Core) (row int, err error)
	GetUser() (user []Core, err error)
	UpdateId(data Core) (row int, err error)
	Upgrade(input Core) (data Core, err error)
}
