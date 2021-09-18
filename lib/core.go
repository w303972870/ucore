package lib
import(
    "unsafe"
    "runtime"
)

type Us struct{}
var Usys Us 

/*获取程序占用的内存*/
func ( t * Us ) MemOfPro () uint64 {
    var mem runtime.MemStats
    runtime.ReadMemStats( &mem )
    return mem.Alloc
}

/*获取变量占用的内存*/
func ( t * Us ) MemOfParm ( p interface{} ) uintptr {
    return unsafe.Sizeof( p )
}

