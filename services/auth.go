package services

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/realjv3/gotasks/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) domain.AuthService {
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
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(os.Getenv("JWT_KEY"))
}

func generatePrivateKey() *ecdsa.PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Error generating private key - %v", err)
	}

	return key
}

func savePrivateKey(key *ecdsa.PrivateKey) error {
	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		log.Fatalf("Error marshalling private key - %v", err)
	}

	keyPEM := pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyBytes,
	}

	file, err := os.Create("private.pem")
	if err != nil {
		log.Fatalf("Error creating private.pem - %v", err)
	}
	defer file.Close()

	return pem.Encode(file, &keyPEM)
}

func loadPrivateKey() (*ecdsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	
	return key, nil
}
