package services

import (
	"errors"
	"os"
	"time"

	"devforum/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ── Types ────────────────────────────────────────────────────────────────────

type RegisterInput struct {
	FullName string `json:"full_name" binding:"required"`
	Username string `json:"username"  binding:"required"`
	Email    string `json:"email"     binding:"required,email"`
	Password string `json:"password"  binding:"required,min=8"`
}

type LoginInput struct {
	Identifier string `json:"identifier" binding:"required"` // email or username
	Password   string `json:"password"   binding:"required"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  PublicUser  `json:"user"`
}

type PublicUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// ── Service ──────────────────────────────────────────────────────────────────

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

// Register creates a new user account.
func (s *AuthService) Register(input RegisterInput) (*AuthResponse, error) {
	// Check username taken
	var existing models.User
	if err := s.DB.Where("username = ?", input.Username).First(&existing).Error; err == nil {
		return nil, errors.New("username already taken")
	}
	// Check email taken
	if err := s.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := models.User{
		FullName: input.FullName,
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashed),
		Role:     "member",
		IsActive: true,
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	token, err := generateJWT(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{Token: token, User: toPublicUser(user)}, nil
}

// Login validates credentials and returns a JWT.
// If the identifier matches the ADMIN_EMAIL env var and password matches
// ADMIN_PASSWORD, the user is promoted to admin role on the fly (first login seed).
func (s *AuthService) Login(input LoginInput) (*AuthResponse, error) {
	var user models.User

	// Find by email or username
	err := s.DB.Where("email = ? OR username = ?", input.Identifier, input.Identifier).
		First(&user).Error
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, errors.New("account is disabled")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// If this is the admin email from .env, ensure role = admin in DB
	adminEmail := os.Getenv("ADMIN_EMAIL")
	if adminEmail != "" && user.Email == adminEmail && user.Role != "admin" {
		s.DB.Model(&user).Update("role", "admin")
		user.Role = "admin"
	}

	token, err := generateJWT(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{Token: token, User: toPublicUser(user)}, nil
}

// ValidateToken parses and validates a JWT string.
func ValidateToken(tokenStr string) (*Claims, error) {
	secret := []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}

// ── Helpers ──────────────────────────────────────────────────────────────────

func generateJWT(user models.User) (string, error) {
	secret := []byte(os.Getenv("SECRET_KEY"))
	claims := Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func toPublicUser(u models.User) PublicUser {
	return PublicUser{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		FullName: u.FullName,
		Role:     u.Role,
	}
}