package httpx

// UserAgent mirrors the official OpenCode CLI identity string:
// opencode/<channel>/<version>/<name>. Stable releases use channel "latest".
func UserAgent() string {
	return "opencode/latest/1.19.0/cli"
}
