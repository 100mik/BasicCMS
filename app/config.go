package app

func ConfigureInstance() {
	Configuration = Settings{
		ServerAddress: ":9000",
		DBUsername: "CMSUser",
		DBPassword: "mycms",
		DBName:     "basicCMS",
	}
}

