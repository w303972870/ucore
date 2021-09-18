package ulog
import(
    "fmt"
    "log"
	"os"
	"path"
	"time"
    "runtime"
    "ucore/lib"
    "github.com/sirupsen/logrus"
    "github.com/gin-gonic/gin"
)
type HtLogs struct{}

/*公用日志*/
var ULogs HtLogs

func( mt * HtLogs ) Waring ( message interface{} ) {
    if lib.LibConfigParms.SysType == "windows" {
        fmt.Println("[警告]", message.(string) )
    } else {
        fmt.Printf("%c[7;46;33m[警告]%s%c[0m\n", 0x1B, message.(string), 0x1B)
    }
}

func( mt * HtLogs ) Error ( message interface{} ) {
    if lib.LibConfigParms.SysType == "windows" {
        fmt.Println("[错误]", message.(string) )
    } else {
        fmt.Printf("%c[5;41;32m[错误]%s%c[0m\n", 0x1B, message.(string), 0x1B)
    }
    os.Exit( 1 )
}

func( mt * HtLogs ) Sys ( message interface{} ) {
    if lib.LibConfigParms.SysType == "windows" {
        fmt.Println("[系统]", message.(string) )
    } else {
        fmt.Printf("%c[1;40;32m[系统]%s%c[0m\n", 0x1B, message.(string), 0x1B)
    }
}

/*将接口的返回统一格式化*/
func( mt * HtLogs ) GinJson ( code int , msg string , data interface{} ) ( map[string]interface{} ) {
    return gin.H{"code": code , "msg" : msg , "data" : data}
}

func( mt * HtLogs ) Info ( message interface{} ) {
    fmt.Println( "[信息]" , message.(string) )
}

func( mt * HtLogs ) FInfoLog ( file string , message string ) {
   	logFile, err := os.OpenFile( file , os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )
	if nil != err {
		mt.Error( err.Error() )
	}
	loger := log.New(logFile, "", log.Ldate|log.Ltime )
	loger.SetFlags(log.Ldate | log.Ltime )
    loger.Println( "[" , runtime.NumGoroutine() , "]" , message )
	if err := logFile.Close(); err != nil {
		mt.Error( err )
	}
}

/*请求日志*/
func( mt * HtLogs ) AccessLogHandle() gin.HandlerFunc {

	fileName := path.Join( lib.LibConfigParms.Config.LogsDir , lib.LibConfigParms.Config.LogsName )
	src, err := os.OpenFile( fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend )

	if err != nil {
		ULogs.Waring( err )
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel( logrus.InfoLevel )
    logger.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat:"2006-01-02 15:04:05",
    })
    //logger.SetOutput(os.Stdout)
    

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()
        
        // 日志格式
		logger.WithFields(logrus.Fields{
			"status"  : statusCode,
			"latency_time" : latencyTime,
			"ip"           : clientIP,
			"m"            : lib.Usys.MemOfPro() / 1024 ,
			"method"       : reqMethod,
			"uri"          : reqUri,
            "referer"      : c.Request.Referer() ,
			"agent"        : c.Request.UserAgent(),
		}).Info()
	}
}