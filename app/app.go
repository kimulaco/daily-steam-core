package app

type App struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	Title      string `json:"title"`
	ThumbUrl   string `json:"thumbnail_url"`
	ReleasedAt string `json:"released_at"`
	Price      string `json:"price"`
	SalePrice  string `json:"sale_price"`
}

func (a App) IsValid() bool {
	return (a.Id != "" &&
		a.Url != "" &&
		a.Title != "" &&
		a.ThumbUrl != "" &&
		a.ReleasedAt != "" &&
		a.Price != "")
}
