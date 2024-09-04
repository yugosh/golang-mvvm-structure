package mappers

import (
	"BACKEND-GOLANG-MVVM/internal/app/models"
	"BACKEND-GOLANG-MVVM/internal/app/viewmodels"
)

func ToUserViewModel(user *models.User) *viewmodels.UserViewModel {
	return viewmodels.NewUserViewModel(user.ID, user.Username, user.Email)
}

// Di file internal/app/mappers/user_mapper.go, kita akan membuat fungsi untuk memetakan User ke UserViewModel.
