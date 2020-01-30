package route

import (
	authController "gin-skeleton/app/controllers/auth"
	WebsocketController "gin-skeleton/app/controllers/websocket"
	"gin-skeleton/app/models"
	_ "gin-skeleton/app/validators"
	"gin-skeleton/database"
	"gin-skeleton/helper/render"
	"github.com/gin-contrib/cors"
	"io"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	db          = database.DB
	identityKey = "id"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 	设备日志文件
	f, _ := os.Create("runtime.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// 	CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = false
	config.AddAllowHeaders("authorization")
	router.Use(cors.New(config))

	// 	JWT
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "gin-skeleton",
		Key:           []byte(os.Getenv("JWT_SECRET")),
		Timeout:       5 * time.Hour,
		MaxRefresh:    7 * 24 * time.Hour,
		IdentityKey:   identityKey,
		TokenLookup:   "header: Authorization, query: token, cookie: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		// 	验证动作
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var user models.User

			if err := c.ShouldBind(&user); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			db.Where("account = ?", user.Account).Find(&user)
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &user, nil
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			render.Data(c, map[string]string{"token": token})
		},
	})

	{

		api := router.Group("api")
		api.Use(authMiddleware.MiddlewareFunc())

		router.POST("/api/auth/login", authMiddleware.LoginHandler)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.GET("/refresh_token", authMiddleware.RefreshHandler)
			auth.GET("/info", authController.Info)
		}

		websocket := api.Group("ws")
		{
			websocket.GET("", WebsocketController.Websocket)
		}
	}

	return router
}
