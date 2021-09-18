package main
import (
    "ucore/lib"
    "ucore/common"
)
/*
#include <unistd.h>
*/
import "C"
func main(){
    if lib.LibConfigParms.Config.Daemon == true {
        C.daemon(1, 1)
    }
    ucore.Start()
}

func init() {
    ucore.InitSystem()
}
