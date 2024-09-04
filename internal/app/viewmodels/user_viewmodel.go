// Di file internal/app/viewmodels/user_viewmodel.go, kita akan mendefinisikan struktur yang akan dikirim sebagai respons.

package viewmodels

type UserViewModel struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserViewModel(id uint, username, email string) *UserViewModel {
	return &UserViewModel{
		ID:       id,
		Username: username,
		Email:    email,
	}
}
