package tables

import "github.com/NoAnguish/PearlerBackend/backend/utils/database"

func MigrBaseSubscriptions(s *database.Session) {
	query := `CREATE TABLE "Subscriptions" (
		"id" VARCHAR(18) PRIMARY KEY,
		"source" VARCHAR(18),
		"target" VARCHAR(18),
		"deleted" BOOLEAN
	);
	CREATE INDEX "IdxSubscriptions" ON "Subscriptions" (
		"source",
		"target"
	);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}
