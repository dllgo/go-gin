package ginplus

// func main() {
// 	mconf := gins.Config{Address: ":8600", ReadTimeout: 30, WriteTimeout: 30}
// 	httpserver, err := gins.NewServerHttp(mconf)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	httpserver.Router = InitRouter()
// 	err = httpserver.Listen()
// 	defer httpserver.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }
// func InitRouter() *gin.Engine {
// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	router.NoRoute(func(c *gin.Context) {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"code": 404,
// 			"msg":  "找不到该路由",
// 		})
// 		return
// 	})
// 	router.NoMethod(func(c *gin.Context) {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"code": 404,
// 			"msg":  "找不到该方法",
// 		})
// 		return
// 	})
// 	//注册api
// 	registerApiRouter(router)
// 	return router
// }
// func registerApiRouter(router *gin.Engine) {
// 	// user.NewUserHandler().Router(router)
// }
