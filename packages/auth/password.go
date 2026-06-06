package auth


func HashPassword(password string) string {
	return "hashed-" + password
}
