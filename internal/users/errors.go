package users

type WrongUsernameOrPasswordError struct{}

func (w *WrongUsernameOrPasswordError) Error() string {
	return "incorrect username or password"
}
