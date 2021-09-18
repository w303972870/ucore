package ucore

import(
    "github.com/gin-gonic/gin"
    "ucore/lib"
    "ucore/logs"
)

func InitGin(){
    //gin.DisableConsoleColor()
    //gin.SetMode( gin.ReleaseMode )
    
    lib.LibConfigParms.SysGin.Use( ulog.ULogs.AccessLogHandle() )
    lib.LibConfigParms.SysGin.Use( gin.Recovery() )
    lib.LibConfigParms.SysGin.Use( checkIsOk() )
    InitRoute()
}

/*检查合法性，暂时内网使用，先不加密钥验证*/
func checkIsOk() gin.HandlerFunc { return func( c *gin.Context ) {

    platform := c.Param("platform")

    if UTools.StrInSlice( platform , lib.LibConfigParms.Platform ) && UTools.StrInSlice( c.ClientIP() , lib.LibConfigParms.WhiteIps ) {
        lib.LibConfigParms.RequestPlatform = platform
        c.Next()
    } else {
        c.JSON( lib.StatusServiceUnavailable , ulog.ULogs.GinJson( lib.StatusServiceUnavailable , "非法请求", nil ) )
    }
}}
