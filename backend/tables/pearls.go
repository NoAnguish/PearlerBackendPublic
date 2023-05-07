package tables

import "github.com/NoAnguish/PearlerBackend/backend/utils/database"

func MigrPearls(s *database.Session) {
	query := `CREATE TABLE "Pearls" (
		"id" VARCHAR(18) PRIMARY KEY,
		"account_id" VARCHAR(18),
		"cocktail_id" VARCHAR(18),
		"grade" INT,
		"review" TEXT,
		"created_at" BIGINT
	);
	CREATE INDEX "IdxPearls" ON "Pearls" (
		"id", "created_at" DESC
	);
	CREATE INDEX "IdxPearlsAccount" ON "Pearls" (
		"account_id", "created_at" DESC
	);
	CREATE INDEX "IdxPearlsCocktail" ON "Pearls" (
		"cocktail_id", "created_at" DESC
	);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrPearlsAddImages(s *database.Session) {
	query := `
		ALTER TABLE "Pearls" ADD COLUMN "image_id" VARCHAR(18);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrPearlsFixImages(s *database.Session) {
	query := `
		ALTER TABLE "Pearls" DROP COLUMN "image_id";
		ALTER TABLE "Pearls" ADD COLUMN "image_id" VARCHAR(18) DEFAULT '';
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrPearlsAddImageURLs(s *database.Session) {
	query := `
		ALTER TABLE "Pearls" DROP COLUMN "image_id";
		ALTER TABLE "Pearls" ADD COLUMN "image_url" VARCHAR(100) DEFAULT '';
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}
