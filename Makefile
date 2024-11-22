# Đường dẫn tới file migration.go
MIGRATION_FILE=database/migrations/migration.go

# Lệnh Migration Up
migrateup:
	go run $(MIGRATION_FILE)