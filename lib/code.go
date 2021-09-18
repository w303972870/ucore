package lib

const (
    StatusOK                   = 200 // RFC 7231, 6.3.1
    StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
    StatusBadRequest                   = 400 // RFC 7231, 6.5.1
    StatusUnauthorized                 = 401 // RFC 7235, 3.1
    StatusForbidden                    = 403 // RFC 7231, 6.5.3
    StatusNotFound                     = 404 // RFC 7231, 6.5.4
    StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12

    StatusInternalServerError           = 500 // RFC 7231, 6.6.1
    StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
    StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
)