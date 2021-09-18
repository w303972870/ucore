package ucore
import (
    "strconv"
    "ucore/lib"
)

func Start() {
    lib.LibConfigParms.SysGin.Run( UTools.Str( lib.LibConfigParms.Config.Host , ":" , strconv.Itoa(lib.LibConfigParms.Config.Port ) ) )
}
