package view

import (
	"fmt"
	"strings"

	"github.com/NoAnguish/PearlerBackend/backend/objects/account"
	"github.com/NoAnguish/PearlerBackend/backend/objects/cocktail"
	"github.com/NoAnguish/PearlerBackend/backend/objects/pearl"
	"github.com/NoAnguish/PearlerBackend/backend/objects/subscription"
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
)

func GetByNamePrefix(s *database.Session, accountId string, prefix string, limit uint) (*[]account.Account, *api_errors.Error) {
	lowerPrefix := strings.ToLower(prefix) + "%"
	query := `
		SELECT
			*
		FROM (
			SELECT
				*
			FROM "%s"
			WHERE id NOT IN (
				SELECT 
					target AS id
				FROM "%s" 
				WHERE source = '%s' AND NOT deleted
			)
		) filled_users
		WHERE (id != '%s') AND (LOWER(name) LIKE '%s')
		ORDER BY (name)
		LIMIT %d;
	`
	query = fmt.Sprintf(query, account.TableName, subscription.TableName, accountId, accountId, lowerPrefix, limit)
	data, err := database.Get[account.Account](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}

func GetPearlEvents(s *database.Session, accountId string, limit uint) (*[]PearlView, *api_errors.Error) {
	query := `
		SELECT 
			pearls_n_acc.*,
			cocktails.name AS cocktail_name,
			cocktails.image_url AS cocktail_image_url
		FROM (
			SELECT
				pearls.*,
				accounts.name AS account_name,
				accounts.image_url AS account_image_url
			FROM (
				SELECT
					id, account_id, cocktail_id, review, grade, created_at, image_url
				FROM "%s"
				WHERE account_id IN (
					SELECT 
						target AS account_id
					FROM "%s" 
					WHERE source = '%s' AND NOT deleted

					UNION
					
					SELECT '%s' AS account_id
				)
			) pearls

			LEFT JOIN "%s" accounts
			ON (pearls.account_id = accounts.id)
		) pearls_n_acc

		LEFT JOIN "%s" cocktails
		ON (pearls_n_acc.cocktail_id = cocktails.id)
		ORDER BY (created_at) DESC
		LIMIT %d;
	`
	query = fmt.Sprintf(query, pearl.TableName, subscription.TableName, accountId, accountId, account.TableName, cocktail.TableName, limit)
	data, err := database.Get[PearlView](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}

func GetAccountSubscriptions(s *database.Session, accountId string) (*[]account.Account, *api_errors.Error) {
	query := `
		SELECT
			*
		FROM "%s"
		WHERE id IN (
			SELECT 
				target AS id
			FROM "%s" 
			WHERE source = '%s' AND NOT deleted
		)
		ORDER BY (name);
	`
	query = fmt.Sprintf(query, account.TableName, subscription.TableName, accountId)
	data, err := database.Get[account.Account](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}

func GetFilledPearlsByCocktailId(s *database.Session, cocktailId string) (*[]PearlView, *api_errors.Error) {
	query := `
		SELECT 
			pearls_n_acc.*,
			cocktails.name AS cocktail_name,
			cocktails.image_url AS cocktail_image_url
		FROM (
			SELECT
				pearls.*,
				accounts.name AS account_name,
				accounts.image_url AS account_image_url
			FROM (
				SELECT
					id, account_id, cocktail_id, review, grade, created_at, image_url
				FROM "%s"
				WHERE cocktail_id = '%s' 
			) pearls

			LEFT JOIN "%s" accounts
			ON (pearls.account_id = accounts.id)
		) pearls_n_acc

		LEFT JOIN "%s" cocktails
		ON (pearls_n_acc.cocktail_id = cocktails.id)
		ORDER BY created_at DESC;
	`
	query = fmt.Sprintf(query, pearl.TableName, cocktailId, account.TableName, cocktail.TableName)
	data, err := database.Get[PearlView](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}

func GetFilledPearlsByAccountId(s *database.Session, accountId string) (*[]PearlView, *api_errors.Error) {
	query := `
		SELECT 
			pearls_n_acc.*,
			cocktails.name AS cocktail_name,
			cocktails.image_url AS cocktail_image_url
		FROM (
			SELECT
				pearls.*,
				accounts.name AS account_name,
				accounts.image_url AS account_image_url
			FROM (
				SELECT
					id, account_id, cocktail_id, review, grade, created_at, image_url
				FROM "%s"
				WHERE account_id = '%s' 
			) pearls

			LEFT JOIN "%s" accounts
			ON (pearls.account_id = accounts.id)
		) pearls_n_acc

		LEFT JOIN "%s" cocktails
		ON (pearls_n_acc.cocktail_id = cocktails.id)
		ORDER BY created_at DESC;
	`
	query = fmt.Sprintf(query, pearl.TableName, accountId, account.TableName, cocktail.TableName)
	data, err := database.Get[PearlView](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}
