package router

import "github.com/gin-gonic/gin"

func CombineRouter(r *gin.Engine) {
	authRoute(r.Group("/auth"))
	pinRoute(r.Group("/auth/pin"))
	userRoute(r.Group("/user"))
	balanceRouter(r.Group("/balance"))
	transactionRoute(r.Group("/transaction"))
	historyRoute(r.Group("/history"))
}
