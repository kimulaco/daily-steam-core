package app

import (
	"database/sql"
	"log"
)

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
