package bootstrap

func Initialize() {
	startupVariables()
	ConnectToDB()
	migrateDB()
}
