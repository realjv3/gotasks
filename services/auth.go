package services

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/realjv3/gotasks/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	keyFileName = "private.pem"
	keyType     = "EC PRIVATE KEY"
)

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) domain.AuthService {
	if _, err := os.Stat(keyFileName); err != nil {
		log.Println("Generating private key...")

		key, err := generatePrivateKey()
		if err != nil {
			log.Fatalf("Error generating private key - %v", err)
		}

		err = savePrivateKey(key)
	}

	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Login(userID int, password string) (string, error) {
	user, err := s.userRepo.Get(userID)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": strconv.Itoa(user.ID),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	key, err := loadPrivateKey()
	if err != nil {
		return "", err
	}

	return token.SignedString(key)
}

func generatePrivateKey() (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func savePrivateKey(key *ecdsa.PrivateKey) error {
	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	keyPEM := pem.Block{
		Type:  keyType,
		Bytes: keyBytes,
	}

	file, err := os.Create(keyFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, &keyPEM)
}

func loadPrivateKey() (*ecdsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(keyFileName)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != keyType {
		return nil, err
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
