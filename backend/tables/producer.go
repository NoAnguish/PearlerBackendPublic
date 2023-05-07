package tables

import "github.com/NoAnguish/PearlerBackend/backend/utils/database"

var tableVersions = map[string]func(*database.Session){
	"v_00001": MigrBaseAccount,
	"v_00002": MigrAccountIndexes,
	"v_00003": MigrBaseSubscriptions,
	"v_00004": MigrCocktails,
	"v_00005": MigrPearls,
	"v_00006": MigrCocktailsRemovePearls,
	"v_00007": MigrImages,
	"v_00008": MigrCocktailsAddImages,
	"v_00009": MigrPearlsAddImages,
	"v_00010": MigrAccountAddImages,
	"v_00011": MigrCocktailsFixImages,
	"v_00012": MigrPearlsFixImages,
	"v_00013": MigrAccountFixImages,
	"v_00014": MigrCocktailsFixRecipeDescriptionTypes,
	"v_00015": MigrAccountsAddImageURLs,
	"v_00016": MigrCocktailsAddImageURLs,
	"v_00017": MigrPearlsAddImageURLs,
	"v_00018": MigrImagesDropTable,
	"v_00019": MigrAccountsAddDescription,
	"v_00020": MigrCocktailsChangeImageURLFieldLength,
}

func GetTableData() map[string]func(*database.Session) {
	return tableVersions
}
