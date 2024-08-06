package api

import (
	"api-service/api/handlers"

	pbMedal "github.com/Bekzodbekk/protofiles/genproto/medals"
	pbUser "github.com/Bekzodbekk/protofiles/genproto/user"
	"github.com/gin-gonic/gin"
)

func NewGin(
	userService pbUser.UserServiceClient,
	medalService pbMedal.MedalServiceClient,
) *gin.Engine {
	r := gin.Default()

	hnd := handlers.Handlers{
		UserService:  userService,
		MedalService: medalService,
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", hnd.Register)
		auth.POST("/login", hnd.Login)
		auth.POST("/refresh", hnd.RefreshToken)
	}

	medals := r.Group("/medals")
	{
		medals.GET("/ranking")
		medals.POST("/", hnd.CreateMedal)
		medals.PUT("/:id", hnd.UpdateMedal)
		medals.DELETE("/:id", hnd.DeleteMedal)
	}

	events := r.Group("/events")
	{
		events.GET("/")
		events.GET("/:id")
		events.POST("/")
		events.PUT("/:id")
		events.DELETE("/:id")
	}

	athletes := r.Group("/athletes")
	{
		athletes.GET("/")
		athletes.GET("/:id")
		athletes.POST("/")
		athletes.PUT("/:id")
		athletes.DELETE("/:id")
	}

	live := r.Group("/live")
	{
		live.GET("/:eventId")
		live.POST("/:eventId")
	}
	return r
}
