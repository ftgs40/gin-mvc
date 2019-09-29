package configs

func GetServerConfig() (config map[string]string) {
	config = make(map[string]string)

	config["HOST"] = "0.0.0.0"
	config["PORT"] = ":8080"
	return
}
