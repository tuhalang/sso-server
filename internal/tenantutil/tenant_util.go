package tenantutil

import "strings"

const (
	DefaultTenant = "normal"
)

func GetTenant(username string) (string, string) {
	data := strings.Split(username, "/")
	if len(data) == 2 {
		return data[0], data[1]
	}
	return DefaultTenant, username
}
