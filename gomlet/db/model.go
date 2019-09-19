package db

// Role ...
type Role int

const (
	// Admin ...
	Admin Role = 0
	// User ...
	User Role = 10
)

//func (r Role) String() string { return months[(r-1)%12] }

// var roles = [...]string{
// 	"ADMIN",
// 	"USER",
// }

// var Role_value = map[string]int32{
// 	"ADMIN": 0,
// 	"USER":  10,
// }

// UserModel ...
type UserModel struct {
	UserID   string `bson:"user_id" json:"user_id"`
	Password string
	Name     string
	Role     Role
	// SSHKeys  []string
}
