package ucore
/*
   初始化参数
*/
import (
    "os"
    "runtime"
    "syscall"
    "ucore/lib"
    "ucore/logs"
    "path/filepath"
)

func init(){
    dir, _ := os.Executable()
    exPath := filepath.Base(dir)
    if exPath != "ucore" {
        ulog.ULogs.Error( "Command `ucore` Name Error!" )
    }
    setCpus()

    lib.LibConfigParms.SysType = runtime.GOOS
}

func setCpus(){
    var rLimit syscall.Rlimit
    err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
    if err != nil {
        ulog.ULogs.Error( UTools.Str( "获取系统变量ulimit错误： " , err.Error() ) )
    }
    if rLimit.Cur < 65535 {
        rLimit.Max = 204800
        rLimit.Cur = 204800
        err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
        if err != nil {
            ulog.ULogs.Error( UTools.Str( "设置系统变量ulimit错误： " , err.Error() ) )
        }
    }

    if runtime.NumCPU() > 24 {
        lib.LibConfigParms.Cores = 10000
    } else {
        lib.LibConfigParms.Cores = 5000
    }
    runtime.GOMAXPROCS( runtime.NumCPU() )
}

func InitSystem(){
    InitConf()
    InitGin()
}