package ucore
import(
    "fmt"
    "net"
    "os"
    "bytes"
    "unsafe"
    "runtime"
    "path/filepath"
)

type Tools struct{}
var UTools Tools 

/*判断文件或文件夹是否存在，true：存在，false：不存在，err不为空说明有错*/
func( t * Tools )Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/*获取程序占用的内存*/
func ( t * Tools ) MemOfPro () uint64 {
    var mem runtime.MemStats
    runtime.ReadMemStats( &mem )
    return mem.Alloc
}

/*获取变量占用的内存*/
func ( t * Tools ) MemOfParm ( p interface{} ) uintptr {
    return unsafe.Sizeof( p )
}

/*查找一个字符串是否在一个切片中，true：存在，false：不存在*/
func ( t * Tools ) StrInSlice ( a string, list []string ) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

/*拼接字符串*/
func( t * Tools )Str( parms ...string ) string {
    var buffer bytes.Buffer

    for _ , parm := range parms {
        buffer.WriteString( parm )
    }
    return buffer.String()
}

/*退出程序*/
func( t * Tools )Bye( code int ){
    os.Exit( code )
}

/*获取当前cgi bin名称*/
func( t * Tools )BinName() string {
    _, file := filepath.Split( os.Args[0] )
    return file
}

/*创建目录*/
func ( t * Tools )MkDir( tdir string ) {
    if is , _ := UTools.Exist( tdir ) ; !is {
         os.MkdirAll( tdir , os.ModePerm )
    }
}

/*获取变量类型*/
func ( t * Tools )Type( parm interface{} ) string {
    switch t := parm.(type) {
    case bool:
        return "bool"
    case int:
        return "int"
    default:
        return fmt.Sprintf("%T", t)
    }
}

/*判断ip格式*/
func ( t * Tools ) IsIp ( ip string ) bool {
    if addr := net.ParseIP( ip ); addr == nil {
         return false
    } else {
         return true
    }
}

/*创建文件*/
func ( t * Tools ) Touch ( file string ) ( bool , *os.File ) {
    f, err := os.Create( file )
    if err != nil {
        return false , nil
    }
    return true , f
}

