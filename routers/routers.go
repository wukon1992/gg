package routers

import (
	"gg/pkg/setting"
	v1 "gg/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//标签
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags/add", v1.AddTag)
		apiv1.PUT("/tag/edit/:id", v1.EditTag)
		apiv1.DELETE("/tag/delete/:id", v1.DeleteTag)

		//文章
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles:id", v1.EditArticle)
		apiv1.DELETE("/articles:id", v1.DeleteArticle)
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
