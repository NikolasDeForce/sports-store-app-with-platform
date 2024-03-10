package identity

type SignInManager interface {
	SignIn(user User) error
	SignOut(use User) error
}
