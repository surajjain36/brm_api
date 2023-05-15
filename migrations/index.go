package migrations

func PreAutoMigrate() {

}

func Migrate() {
	createDependencyTables()
	createUserRelationsTables()
	createSuperAdmin()
}

func PostMigrate() {

}
