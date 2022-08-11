package app

import (
	"github.com/kimulaco/daily-steam-core/date"
)

type App struct {
	Id         string
	Url        string
	Title      string
	ThumbUrl   string
	ReleasedAt date.Date
	Price      string
	SalePrice  string
}

type AppJson struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	Title      string `json:"title"`
	ThumbUrl   string `json:"thumbnail_url"`
	ReleasedAt string `json:"released_at"`
	Price      string `json:"price"`
	SalePrice  string `json:"sale_price"`
}

func (a App) ToJson() AppJson {
	return AppJson{
		Id:         a.Id,
		Url:        a.Url,
		Title:      a.Title,
		ThumbUrl:   a.ThumbUrl,
		ReleasedAt: a.ReleasedAt.ToString(),
		Price:      a.Price,
		SalePrice:  a.SalePrice,
	}
}
