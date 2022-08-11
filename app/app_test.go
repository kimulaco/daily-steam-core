package app

import (
	"testing"

	"github.com/kimulaco/daily-steam-core/date"
)

func TestApp(t *testing.T) {
	a := App{
		Id:       "test_app_id",
		Url:      "https://github.com/kimulaco/daily-steam-core",
		Title:    "TestApp",
		ThumbUrl: "",
		ReleasedAt: date.Date{
			Year:  2022,
			Month: 1,
			Day:   1,
		},
		Price:     "1000",
		SalePrice: "500",
	}

	t.Run("ToJson", func(t *testing.T) {
		j := a.ToJson()
		if j.ReleasedAt != "2022/1/1" {
			t.Error(j.ReleasedAt)
		}
	})
}
