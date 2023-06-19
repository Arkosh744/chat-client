package model

func BuildRedisRefreshKey(username string) string {
	return "user:" + username + ":refresh"
}

func BuildRedisAccessKey(username string) string {
	return "user:" + username + ":access"
}
