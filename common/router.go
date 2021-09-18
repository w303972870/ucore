package ucore

import (
    "ucore/controller"
    "ucore/lib"
    "github.com/gin-gonic/gin"
)

func InitRoute() {
    v1 := lib.LibConfigParms.SysGin.Group("/v1")
    initV1( v1 )
}

/*V1版本的api*/
func initV1( v1 *gin.RouterGroup ){
	v1.POST("/:platform/:cn/auth", controller.Auth ) /*登录*/
	v1.POST("/:platform/group/all", controller.AllGroups ) /*获取所有部门*/
	v1.POST("/:platform/user/all", controller.AllUsers ) /*获取所有用户*/
	v1.POST("/:platform/group/:gid/", controller.GetGroupById ) /*根据id获取部门*/
	v1.POST("/:platform/user/:uid/", controller.GetUserById ) /*根据id获取用户*/
	v1.POST("/:platform/:cn/user", controller.GetUserByCn ) /*根据登录名获取用户*/
}



