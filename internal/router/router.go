package router

import (
	v1 "MengCDN/internal/api/v1"
	"MengCDN/internal/middleware"
	"MengCDN/internal/service"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("indexF", "web/front/Index/dist/index.html")
	r.AddFromFiles("indexF3", "web/front/Index/dist/404.html")
	r.AddFromFiles("admin", "web/front/MengCDN_Admin/dist/index.html")
	r.AddFromFiles("gh", "web/front/MengCDN_Front/dist/index.html")
	return r
}

func InitRouter() {
	router := gin.Default()

	router.HTMLRender = createMyRender()
	router.Use(middleware.CORSMiddleware())

	// 设置静态文件路由
	router.Static("MengCDNAdmin/assets", "web/front/MengCDN_Admin/dist/assets")
	router.Static("assets", "web/front/Index/dist/assets")
	router.Static("browseGH/assets", "web/front/MengCDN_Front/dist/assets")

	// 设置路由
	AuthRouterV1 := router.Group("api/v1")
	AuthRouterV1.Use(middleware.JwtToken())
	{
		// 编辑白名单
		AuthRouterV1.PUT("/cdnWL/:mk", v1.PUTCdnWL)
		// 编辑CDN模块开关
		AuthRouterV1.PUT("/cdnSW/:id/:mk/:sw", v1.PUTCdnSW)
	}
	PublicRouterV1 := router.Group("api/v1")
	{
		//查询白名单
		PublicRouterV1.POST("/cdnWL/:mk", v1.CdnWL)
		//查询CDN模块开关
		PublicRouterV1.POST("/cdnSW/:id/:mk", v1.CdnSW)
		// 登录
		PublicRouterV1.POST("login", v1.Login)
	}

	// Github相关路由
	githubProxyRouter := router.Group("/gh")
	{
		githubProxyRouter.GET("/:owner/:repo/*path", service.GithubProxy)
	}

	// NPM相关路由
	npmProxyRouter := router.Group("/npm")
	{
		npmProxyRouter.GET("/:package/:version/*path", service.NpmProxy)
		npmProxyRouter.GET("/browse/:package/:version/*path", service.NpmProxyBrowse)
	}

	// WordPress相关路由
	wpProxyRouter := router.Group("/wp")
	{
		wpProxyRouter.GET("/theme/:package/:version/*path", service.WpProxyTh)
		wpProxyRouter.GET("/plugin/:package/*path", service.WpProxyPl)
	}

	// 首页和其他页面路由
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "indexF", nil)
	})
	router.GET("/404.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "indexF3", nil)
	})

	// MengCDN Admin页面路由
	router.GET("/MengCDNAdmin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin", nil)
	})

	// 错误处理
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "indexF3", nil) // 使用404页面作为默认错误页面
	})

	// 启动服务器
	err := router.Run(":8001")
	if err != nil {
		panic(err)
	}
}
