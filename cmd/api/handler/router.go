package handler

import (
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	// 公共router
	apiRouter := r.Group("/douyin")

	apiRouter.Use(middleware.Cors())
	apiRouter.POST("/user/register/", Register)
	apiRouter.POST("/user/login/", Login)
	apiRouter.GET("/feed/", Feed)

	apiRouter.Use(middleware.JWT())
	apiRouter.GET("/user/", UserInfo)

	apiRouter.POST("/publish/action/", Publish)
	apiRouter.GET("/publish/list/", PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", FavoriteAction)
	apiRouter.GET("/favorite/list/", FavoriteList)
	apiRouter.POST("/comment/action/", CommentAction)
	apiRouter.GET("/comment/list/", CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", RelationAction)
	apiRouter.GET("/relation/follow/list/", FollowList)
	apiRouter.GET("/relation/follower/list/", FollowerList)
}
