package config

import "time"

var JwtTokenHeader = "authentication"
var JwtContextKey = "jwt_token"
var JwtCookieExpiry = 72 * time.Hour
