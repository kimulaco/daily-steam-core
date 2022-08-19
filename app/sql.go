package app

import (
	"database/sql"
	"log"
	"strconv"
)

type SelectConfig struct {
	Limit int
}

func (c SelectConfig) Init() SelectConfig {
	if c.Limit <= 0 {
		c.Limit = 20
	}
	return c
}

func SelectAppsWhereReleasedAt(
	db *sql.DB,
	releasedAt string,
	c SelectConfig,
) ([]App, error) {
	c = c.Init()
	rows, err := db.Query(`
	SELECT
		id, url, title, thumbnail_url, released_at, price, sale_price
	FROM apps
	WHERE
		released_at = ?
	LIMIT ?
	`, releasedAt, strconv.Itoa(c.Limit))
	if err != nil {
		return []App{}, err
	}

	apps, err := CreateFromSqlRows(rows)
	if err != nil {
		return apps, err
	}

	return apps, nil
}

func CreateFromSqlRows(rows *sql.Rows) ([]App, error) {
	apps := make([]App, 0)

	for rows.Next() {
		var app App
		scanError := rows.Scan(
			&app.Id,
			&app.Url,
			&app.Title,
			&app.ThumbUrl,
			&app.ReleasedAt,
			&app.Price,
			&app.SalePrice,
		)
		if scanError != nil {
			log.Print(scanError)
			continue
		}
		apps = append(apps, app)
	}

	return apps, rows.Err()
}
