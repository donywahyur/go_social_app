package repositories

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"go_social_app/internal/helpers"
	model "go_social_app/internal/models"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterAndInviteUser(user model.User, userInvitation model.UserInvitation) (model.User, error)
	HashPassword(password string) (string, error)
	CompareHash(password, passwordHash string) (bool, error)
	GetUserByID(userID string) (model.User, error)
	ActivationUser(token string) (model.User, error)
	DeleteUser(userID string) error
}

type userRepository struct {
	db    *gorm.DB
	param params
}

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db,
		params{
			memory:      64 * 1024,
			iterations:  3,
			parallelism: 2,
			saltLength:  16,
			keyLength:   32,
		}}
}

func (r *userRepository) registerUser(tx *gorm.DB, user model.User) (model.User, error) {

	err := tx.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
func (r *userRepository) invitationUser(tx *gorm.DB, userInvitation model.UserInvitation) (model.UserInvitation, error) {

	err := tx.Create(&userInvitation).Error
	if err != nil {
		return model.UserInvitation{}, err
	}

	return userInvitation, nil
}
func (r *userRepository) RegisterAndInviteUser(user model.User, userInvitation model.UserInvitation) (model.User, error) {
	err := helpers.RunDBTransaction(r.db, func(tx *gorm.DB) error {
		_, err := r.registerUser(tx, user)
		if err != nil {
			return err
		}

		_, err = r.invitationUser(tx, userInvitation)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return model.User{}, err
	}
	createdUser, err := r.GetUserByID(user.ID)
	if err != nil {
		return model.User{}, err
	}
	return createdUser, nil
}
func (r *userRepository) ActivationUser(token string) (model.User, error) {
	var user model.User
	err := helpers.RunDBTransaction(r.db, func(tx *gorm.DB) error {

		userInvitation, err := r.getUserInvitationByToken(tx, token)
		if err != nil {
			return err
		}

		user, err = r.GetUserByID(userInvitation.UserID)
		if err != nil {
			return err
		}

		user.IsActive = true

		_, err = r.update(tx, user)
		if err != nil {
			return err
		}

		err = r.deleteUserInvitation(tx, userInvitation.UserID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return model.User{}, err
	}

	return user, nil

}
func (r *userRepository) update(tx *gorm.DB, user model.User) (model.User, error) {
	err := tx.Save(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
func (r *userRepository) deleteUserInvitation(tx *gorm.DB, userID string) error {
	err := tx.Where("user_id = ?", userID).Delete(&model.UserInvitation{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) getUserInvitationByToken(tx *gorm.DB, token string) (model.UserInvitation, error) {

	var userInvitation model.UserInvitation

	err := tx.Where("token = ?", token).Where("expired_at > ?", time.Now()).First(&userInvitation).Error

	if err != nil {
		return model.UserInvitation{}, err
	}

	return userInvitation, nil
}

func (r *userRepository) DeleteUser(userID string) error {
	err := helpers.RunDBTransaction(r.db, func(tx *gorm.DB) error {
		err := r.deleteUser(tx, userID)
		if err != nil {
			return err
		}

		err = r.deleteUserInvitation(tx, userID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) deleteUser(tx *gorm.DB, userID string) error {
	err := tx.Where("id = ?", userID).Delete(&model.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByID(userID string) (model.User, error) {
	var user model.User
	err := r.db.Preload("Role").First(&user, "id = ?", userID).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) HashPassword(password string) (string, error) {
	salt := make([]byte, r.param.saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, r.param.iterations, r.param.memory, r.param.parallelism, r.param.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, r.param.memory, r.param.iterations, r.param.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func (r *userRepository) CompareHash(password, passwordHash string) (bool, error) {
	vals := strings.Split(passwordHash, "$")
	if len(vals) != 6 {
		return false, errors.New("invalid hash")
	}

	var memory, time uint32
	var parallelism uint8

	_, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &memory, &time, &parallelism)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return false, err
	}

	decryptedHash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return false, err
	}

	var keyLen = uint32(len(decryptedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, time, memory, parallelism, keyLen)

	return subtle.ConstantTimeCompare(comparisonHash, decryptedHash) == 1, nil

}
