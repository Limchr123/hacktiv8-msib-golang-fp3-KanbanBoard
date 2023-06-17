package entity

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kanban_board/pkg/errs"
	"os"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	FullName  string    `gorm:"column:full_name;not null"`
	Email     string    `gorm:"column:email;uniqueIndex;not null"`
	Password  string    `gorm:"column:password;not null"`
	Role      string    `gorm:"column:role;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Tasks     []Task    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return errs.NewInternalServerError("Error occurred because invalid email format")
	}

	if len(strings.TrimSpace(u.Password)) < 6 {
		return errs.NewInternalServerError("Error occurred because password must 6 letters")
	}

	if u.Role != "admin" && u.Role != "member" {
		return errs.NewInternalServerError("Error occurred because role is unidentified")
	}

	return nil
}

func (u *User) parseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnathenticationError("Error occurred because token is invalid")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, errs.NewUnathenticationError("Error occurred because token is invalid")
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	if id, ok := claim["id"].(float64); !ok {
		return errs.NewUnathenticationError("Error occurred because token is invalid")
	} else {
		u.ID = uint(id)
	}

	if email, ok := claim["email"].(string); !ok {
		return errs.NewUnathenticationError("Error occurred because email is invalid")
	} else {
		u.Email = email
	}

	if role, ok := claim["role"].(string); !ok {
		return errs.NewUnauthorizedError("Error occurred")
	} else {
		u.Role = role
	}

	return nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return errs.NewUnathenticationError("Error occurred because bearer token is invalid")
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return errs.NewUnathenticationError("Error occurred because bearer token is invalid")
	}

	tokenString := splitToken[1]

	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errs.NewUnathenticationError("Error occurred because token is invalid")
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":      u.ID,
		"email":   u.Email,
		"role":    u.Role,
		"expired": time.Now().Add(time.Hour * 2).Unix(),
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}

func (u *User) HashPassword() errs.MessageErr {
	salt := 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)

	if err != nil {
		return errs.NewInternalServerError("Error occurred while trying to hash password")
	}

	u.Password = string(bs)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
