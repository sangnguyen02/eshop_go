package service

import (
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func UserServiceFactory() *UserService {
	return &UserService{repo: repository.NewUserRepository()}
}

// Register creates a new user with hashed password
func (s *UserService) Register(user *model.User, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.UserCredentials = model.UserCredentials{
		Password: string(hashedPassword),
	}
	return s.repo.Create(user)
}

// Login verifies user credentials
func (s *UserService) Login(username, password string) (*model.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.UserCredentials.Password), []byte(password)); err != nil {
		return nil, nil
	}
	return user, nil
}

// SearchUsers searches users by name
func (s *UserService) SearchUsers(name string, page, pageSize int) ([]model.User, int64, error) {
	return s.repo.SearchUsers(name, page, pageSize)
}

// GetUserByID gets a user by ID
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

// UpdatePassword updates the user's password
func (s *UserService) UpdatePassword(userID uint, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	return s.repo.UpdatePassword(userID, string(hashedPassword))
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

// CheckDuplicate checks for duplicate username, phone, or email
func (s *UserService) CheckDuplicate(username, phone, email string) (bool, string, error) {
	user, err := s.repo.CheckDuplicate(username, phone, email)
	if err != nil {
		return false, "", err
	}
	if user != nil {
		if user.Username == username {
			return true, "username", nil
		}
		if user.Phone == phone {
			return true, "phone", nil
		}
		if user.Email == email {
			return true, "email", nil
		}
	}
	return false, "", nil
}

// CheckDuplicateForUpdate checks for duplicate phone or email for update
func (s *UserService) CheckDuplicateForUpdate(id uint, phone, email string) (bool, string, error) {
	user, err := s.repo.CheckDuplicateForUpdate(id, phone, email)
	if err != nil {
		return false, "", err
	}
	if user != nil {
		if user.Phone == phone {
			return true, "phone", nil
		}
		if user.Email == email {
			return true, "email", nil
		}
	}
	return false, "", nil
}
