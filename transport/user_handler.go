package transport

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sprectza/note-taking-app.git/repository"
	"github.com/sprectza/note-taking-app.git/service"
)

type login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user repository.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.RegisterUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "user created successfully"})
}

func (h *UserHandler) LoginHandler(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	user, err := h.userService.LoginUser(loginVals.Email, loginVals.Password)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return &user, nil
}

func (h *UserHandler) AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*repository.User); ok {
				return jwt.MapClaims{
					"id": v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &repository.User{
				Model: gorm.Model{ID: uint(claims["id"].(float64))},
			}
		},
		Authenticator: h.LoginHandler,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*repository.User); ok && v.ID != 0 {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"code": code, "message": message})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return authMiddleware, err
}
