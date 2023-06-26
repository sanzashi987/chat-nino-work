package consts

import "time"

var JwtTokenHeader = "authentication"
var JwtTokenContextKey = "jwt_token"
var JwtUserIDContextKey = "jwt_user_id"
var JwtCookieExpiry = 72 * time.Hour
