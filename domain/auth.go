package domain

type AuthService interface {
	Login(userID int, password string) (string, error) // returns JWT
}
