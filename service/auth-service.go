package service

import (
	"log"

	"github.com/golang_api/dto"
	"github.com/golang_api/entity"
	"github.com/golang_api/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

// AuthService is a what can this service can do
type AuthService interface {
	VerifyCredential(username string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByUsername(username string) entity.User
	IsDuplicateUsername(username string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

// NewAuthService is make a new instance of AuthService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(username string, password string) interface{} {
	res := service.userRepository.VerifyCredential(username, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByUsername(username string) entity.User {
	return service.userRepository.FindByUsername(username)
}

func (service *authService) IsDuplicateUsername(username string) bool {
	res := service.userRepository.IsDuplicateUsername(username)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
