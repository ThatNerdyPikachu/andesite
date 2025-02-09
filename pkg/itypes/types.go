package itypes

import "net/http"

//
type UserRow struct {
	ID        int64 `json:"id"`
	IDS       string
	Snowflake string `json:"snowflake" sqlite:"text"`
	Admin     bool   `json:"admin" sqlite:"tinyint(1)"`
	Name      string `json:"name" sqlite:"text"`
	JoinedOn  string `json:"joined_on" sqlite:"text"`
	PassKey   string `json:"passkey" sqlite:"text"`
	Provider  string `json:"provider" sqlite:"text"`
}

//
type UserAccessRow struct {
	ID   int64  `json:"id"`
	User int64  `json:"user" sqlite:"int"`
	Path string `json:"path" sqlite:"text"`
}

//
type ShareRow struct {
	ID   int64  `json:"id"`
	Hash string `json:"hash" sqlite:"text"` // character(32)
	Path string `json:"path" sqlite:"text"`
}

//
type DiscordRoleAccessRow struct {
	ID        int64  `json:"id"`
	GuildID   string `json:"guild_snowflake" sqlite:"text"`
	RoleID    string `json:"role_snowflake" sqlite:"text"`
	Path      string `json:"path" sqlite:"text"`
	GuildName string `json:"guild_name" sqlite:"text"`
	RoleName  string `json:"role_name" sqlite:"text"`
}

// Middleware provides a convenient mechanism for augmenting HTTP requests
// entering the application. It returns a new handler which may perform various
// operations and should finish by calling the next HTTP handler.
//
// @from https://gist.github.com/gbbr/dc731df098276f1a135b343bf5f2534a
type Middleware func(next http.HandlerFunc) http.HandlerFunc
