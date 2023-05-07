package tables

import "github.com/NoAnguish/PearlerBackend/backend/utils/database"

func MigrImages(s *database.Session) {
	query := `CREATE TABLE "Images" (
		"id" VARCHAR(18) PRIMARY KEY,
		"data" bytea
	);
	CREATE INDEX "IdxImages" ON "Images" (
		"id"
	);
	`
	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}

func MigrImagesDropTable(s *database.Session) {
	query := `DROP TABLE "Images";`

	err := database.Modify(query, s)
	if err != nil {
		panic(err)
	}
}
