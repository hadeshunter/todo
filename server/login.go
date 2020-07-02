package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/hadeshunter/todo/models"
	"github.com/nyaruka/phonenumbers"
)

// LoginResponse contains the client information after successful login
type LoginResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (server *Server) getAuthUser(r *http.Request) (*models.User, error) {
	phone := r.Context().Value(RequestContextKey("phone")).(string)
	return server.db.GetUserByPhone(phone)
}

func (server *Server) login(phone string) (*LoginResponse, error) {
	user, err := server.db.GetUserByPhone(phone)
	if err == gorm.ErrRecordNotFound {
	// 	user, err = server.db.CreateUser(phone)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// } else if err != nil {
		return nil, err
	}
	return &LoginResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Phone: user.Phone,
		Token: getToken(user),
	}, nil
}

func getToken(user *models.User) string {
	fmt.Println(user.Phone)
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Token{Phone: user.Phone})
	signedToken, _ := token.SignedString([]byte(os.Getenv("SERVER_SECRET")))
	return signedToken
}

func (server *Server) formatPhoneNumber(phone string) (string, error) {
	if p, err := phonenumbers.Parse(phone, "VN"); err != nil {
		return "", err
	} else if phonenumbers.IsValidNumber(p) {
		return "0" + strconv.FormatUint(p.GetNationalNumber(), 10), nil
	}
	return "", errors.New("INVALID_PHONE")
}

func (server *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	if phone, err := server.formatPhoneNumber(r.URL.Query().Get("phone")); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else if user, err := server.login(phone); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusAccepted, user)
	}
}
