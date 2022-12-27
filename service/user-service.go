package service

import (
	"log"

	"github.com/golang_api/dto"
	"github.com/golang_api/entity"
	"github.com/golang_api/repository"
	"github.com/mashingan/smapping"
)

// UserService is what can this service can do
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	Delete(userID int64) int64
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) Delete(userID int64) int64 {
	return service.userRepository.DeleteUser(userID)
}
