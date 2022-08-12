package presenter

type User struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Address   string   `json:"address"`
	Phone     string   `json:"phone"`
	Email     string   `json:"email"`
	Latitude  float32  `json:"latitude"`
	Longitude float32  `json:"longitude"`
	Roles     []string `json:"roles"`
	Password  string   `json:"password,omitempty"`
}
