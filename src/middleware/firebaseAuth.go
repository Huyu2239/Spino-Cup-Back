package middleware

import (
	"api/model"
	"api/usecase"
	"context"
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App

func InitFirebase() error {

	serviceAccountKey, _ := base64.StdEncoding.DecodeString(os.Getenv("FB_SERVICE_ACCOUNT_KEY"))

	opt := option.WithCredentialsJSON(serviceAccountKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	firebaseApp = app
	return nil
}

func AuthMiddleware(uu usecase.IUserUsecase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// JWTの取得
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid token")
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			ctx := context.Background()
			client, err := firebaseApp.Auth(ctx)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "failed to initialize Firebase Auth client")
			}

			// JWTの検証（Firebase Authを使用）
			token, err := client.VerifyIDToken(ctx, tokenStr)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Invalid token")
			}

			// firebase_uidの取得
			firebaseUID := token.UID

			// データベースでユーザーを確認または登録
			user, err := uu.GetUserByFirebaseUID(firebaseUID)

			if err != nil {
				// ユーザーが存在しない場合、新規登録
				newUser := model.User{
					FirebaseUID: firebaseUID,
					Email:       token.Claims["email"].(string),
					Name:        token.Claims["name"].(string),
				}
				user, err = uu.CreateUser(newUser)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, "User creation failed")
				}
			}

			// ユーザー情報をコンテキストに保存
			c.Set("user", user)

			// 次のハンドラーを呼び出し
			return next(c)
		}
	}
}
