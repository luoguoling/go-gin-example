package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gin-example/middleware/jwt"
	"go-gin-example/middleware/whiteip"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers/api"
	v1 "go-gin-example/routers/api/v1"
	//_ "github.com/EDDYCJY/go-gin-example/docs"
	 _ "go-gin-example/docs"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(whiteip.IPAuthMiddleware()).GET("/auth",api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT(),whiteip.IPAuthMiddleware())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签列表
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles",v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id",v1.GetArticle)
		//新建文章
		apiv1.POST("/articles",v1.AddArticle)
		//修改文章
		apiv1.PUT("/articles/:id",v1.EditArticle)
		//删除文章
		apiv1.DELETE("/articles/:id",v1.DeleteArticle)

	}
	return r
}
