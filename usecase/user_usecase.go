package usecase

import (
	"os"
	"time"

	"github.com/DaigoSugiyama0317/map-memo-app/model"
	"github.com/DaigoSugiyama0317/map-memo-app/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	LogIn(user model.User) (string, error)
	ResetPassword(user model.User) error
}

type userUsecase struct {
	ur repository.IUser_Repository
}

//コンストラクタ関数
func NewUserUsecase(ur repository.IUser_Repository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	//Passwordからハッシュを生成
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	//(使いまわさずに新しい変数を作り、必要なものだけ入れる)
	newUser := model.User{Email: user.Email, Username: user.Username, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID: newUser.ID,
		Username: newUser.Username,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) LogIn(user model.User) (string, error) {
	//Emailから保存されているユーザーの情報を入手
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	//パスワードを比較
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	//Tokenを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp": time.Now().Add(time.Hour*12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//あとで実装
func (uu *userUsecase) ResetPassword(user model.User) error {
	return nil
}