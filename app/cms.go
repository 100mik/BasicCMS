package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
)

const postsPerPage = 10

func (mycms *App) register (w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err!= nil {
		ResponseWriter(false, "User struct could not be decoded", nil, http.StatusBadRequest, w)
		return
	}

	newUser.Password = strings.TrimSpace(newUser.Password)

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ResponseWriter(false, "Error in hash and salt", nil, http.StatusInternalServerError, w)
		return
	}

	newUser.Username = strings.TrimSpace(newUser.Username)
	newUser.Email = strings.TrimSpace(newUser.Email)
	newUser.Password = string(hash)

	tx := mycms.DB.Begin()
	err = tx.Create(&newUser).Error
	if err != nil {
		tx.Rollback()
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "User created", nil, http.StatusOK, w)
}

func (mycms *App) login (w http.ResponseWriter, r *http.Request){
	userData := User{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		ResponseWriter(false, "Could not decode login information", nil, http.StatusInternalServerError, w)
		return
	}
	userData.Password = strings.TrimSpace(userData.Password)
	userData.Email = strings.TrimSpace(userData.Email)
	userData.Username = strings.TrimSpace(userData.Username)

	user := User{}
	err = mycms.DB.Where("username = ?", userData.Username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		ResponseWriter(false, "User not registered", nil, http.StatusOK, w)
		return
	} else if err != nil {
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if err != nil {
		ResponseWriter(false, "Incorrect password, cannot log in", nil, http.StatusUnauthorized, w)
		return
	}

	currUser := CurrUser{
		ID:       user.ID,
		Username: user.Username,
	}

	err = SetSession(w, currUser, 86400)
	if err != nil {
		fmt.Println(err)
		ResponseWriter(false, "Error in setting session, user not logged in", nil, http.StatusInternalServerError, w)
		return
	}
	ResponseWriter(true, "User logged in and session set", nil, http.StatusOK, w)

}

func (mycms *App) logout (w http.ResponseWriter, r *http.Request) {
	ClearSession(w)
	ResponseWriter(true, "User logged out", nil, http.StatusOK, w)
}

func (mycms *App) addPost (w http.ResponseWriter, r *http.Request) {
	currUser, err := GetCurrUser(r)
	if err!=nil {
		ResponseWriter(false, "User not logged in", nil, http.StatusUnauthorized, w)
		return
	}

	var newPost Post

	err = json.NewDecoder(r.Body).Decode(&newPost)
	if err!= nil {
		ResponseWriter(false, "Post struct could not be decoded", nil, http.StatusBadRequest, w)
		return
	}

	newPost.User = currUser.ID
	 newPost.Title = strings.TrimSpace(newPost.Title)
	 newPost.Content = strings.TrimSpace(newPost.Content)

	tx := mycms.DB.Begin()
	err = tx.Create(&newPost).Error
	if err != nil {
		tx.Rollback()
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "Post created", nil, http.StatusOK, w)
}

func (mycms *App) getPosts (w http.ResponseWriter, r *http.Request) {
	currUser, err := GetCurrUser(r)
	if err!=nil {
		ResponseWriter(false, "User not logged in", nil, http.StatusUnauthorized, w)
		return
	}

	pgs, ok := r.URL.Query()["page"]
	if !ok || len(pgs[0]) < 1 {
		ResponseWriter(false, "Cant list posts", nil, http.StatusBadRequest, w)
		return
	}

	page, err := strconv.Atoi(pgs[0])
	if err != nil {
		ResponseWriter(false, "Parameters not valid. Cannot list posts.", nil, http.StatusBadRequest, w)
		return
	}

	var questions []Post
	offset := (page - 1) * postsPerPage
	err = mycms.DB.Find(&questions).Where("user = ?",currUser.ID).Offset(offset).Limit(postsPerPage).Error
	if err != nil {
		ResponseWriter(false, "Cannot list posts", nil, http.StatusInternalServerError, w)
		return
	}
	ResponseWriter(true, "List of posts", questions, http.StatusOK, w)

}

func (mycms *App) hello (w http.ResponseWriter, r *http.Request) {
	fmt.Print("hello")
}
