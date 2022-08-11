package app

import (
	"testing"
)

func TestAppIsValid(t *testing.T) {
	t.Run("has all the content", func(t *testing.T) {
		a := App{
			Id:         "test_app_id",
			Url:        "https://github.com/kimulaco/daily-steam-core",
			Title:      "TestApp",
			ThumbUrl:   "https://github.com/kimulaco/daily-steam-core/thumb.png",
			ReleasedAt: "2022/1/1",
			Price:      "1000",
			SalePrice:  "500",
		}
		if !a.IsValid() {
			t.Error("this case is that returns true")
		}
	})

	t.Run("no SalePrice", func(t *testing.T) {
		a := App{
			Id:         "test_app_id",
			Url:        "https://github.com/kimulaco/daily-steam-core",
			Title:      "TestApp",
			ThumbUrl:   "https://github.com/kimulaco/daily-steam-core/thumb.png",
			ReleasedAt: "2022/1/1",
			Price:      "1000",
			SalePrice:  "",
		}
		if !a.IsValid() {
			t.Error("this case is that returns true")
		}
	})

	t.Run("no ThumbUrl", func(t *testing.T) {
		a := App{
			Id:         "test_app_id",
			Url:        "https://github.com/kimulaco/daily-steam-core",
			Title:      "TestApp",
			ThumbUrl:   "",
			ReleasedAt: "2022/1/1",
			Price:      "1000",
			SalePrice:  "500",
		}
		if a.IsValid() {
			t.Error("this case is that returns false")
		}
	})
}
