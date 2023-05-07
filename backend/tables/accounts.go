package tables

import "github.com/NoAnguish/PearlerBackend/backend/utils/database"

func MigrBaseAccount(s *database.Session) {
	query := `CREATE TABLE "UserAccounts" (
		"id" VARCHAR(18) PRIMARY KEY,
		"name" VARCHAR(25),
		"email" VARCHAR(50),
		"firebase_uid" VARCHAR(100)
	);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrAccountIndexes(s *database.Session) {
	query := `CREATE INDEX "IdxUserAccountId" ON "UserAccounts" (
		"id",
		"firebase_uid"
	);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrAccountAddImages(s *database.Session) {
	query := `
		ALTER TABLE "UserAccounts" ADD COLUMN "image_id" VARCHAR(18);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrAccountFixImages(s *database.Session) {
	query := `
		ALTER TABLE "UserAccounts" DROP COLUMN "image_id";
		ALTER TABLE "UserAccounts" ADD COLUMN "image_id" VARCHAR(18) DEFAULT '';
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrAccountsAddImageURLs(s *database.Session) {
	query := `
		ALTER TABLE "UserAccounts" DROP COLUMN "image_id";
		ALTER TABLE "UserAccounts" ADD COLUMN "image_url" VARCHAR(100) DEFAULT '';
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrAccountsAddDescription(s *database.Session) {
	query := `
		ALTER TABLE "UserAccounts" ADD COLUMN "description" TEXT DEFAULT '';
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}
