package tables

import "github.com/NoAnguish/PearlerBackend/backend/utils/database"

func MigrCocktails(s *database.Session) {
	query := `CREATE TABLE "Cocktails" (
		"id" VARCHAR(18) PRIMARY KEY,
		"name" VARCHAR(50),
		"recipe" VARCHAR(200),
		"grades_sum" INT,
		"description" VARCHAR(500),
		"pearls_number" INT
	);
	CREATE INDEX "IdxCocktails" ON "Cocktails" (
		"id"
	);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrCocktailsRemovePearls(s *database.Session) {
	query := `
		ALTER TABLE "Cocktails" DROP COLUMN "grades_sum";
		ALTER TABLE "Cocktails" DROP COLUMN "pearls_number";
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrCocktailsAddImages(s *database.Session) {
	query := `
		ALTER TABLE "Cocktails" ADD COLUMN "image_id" VARCHAR(18);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrCocktailsFixImages(s *database.Session) {
	query := `
		ALTER TABLE "Cocktails" DROP COLUMN "image_id";
		ALTER TABLE "Cocktails" ADD COLUMN "image_id" VARCHAR(18) DEFAULT '';
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrCocktailsFixRecipeDescriptionTypes(s *database.Session) {
	query := `
		ALTER TABLE "Cocktails" DROP COLUMN "recipe";
		ALTER TABLE "Cocktails" DROP COLUMN "description";
		ALTER TABLE "Cocktails" ADD COLUMN "recipe" TEXT DEFAULT '';
		ALTER TABLE "Cocktails" ADD COLUMN "description" TEXT DEFAULT '';
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrCocktailsAddImageURLs(s *database.Session) {
	query := `
		ALTER TABLE "Cocktails" DROP COLUMN "image_id";
		ALTER TABLE "Cocktails" ADD COLUMN "image_url" VARCHAR(100) DEFAULT '';
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrCocktailsChangeImageURLFieldLength(s *database.Session) {
	query := `
		ALTER TABLE "Cocktails" ALTER COLUMN "image_url" TYPE VARCHAR(200);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}
