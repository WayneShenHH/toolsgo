package seed

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/models/entities"
	"github.com/WayneShenHH/toolsgo/repository"
)

// AddPlayTypes 加入預設玩法清單
func AddPlayTypes(repo repository.Repository) {
	playsList := []entities.PlayType{
		entities.PlayType{Name: "Point", Code: "point", Description: "讓分"},
		entities.PlayType{Name: "Over Under", Code: "ou", Description: "大小"},
		entities.PlayType{Name: "Money Line", Code: "ml", Description: "獨贏"},
		entities.PlayType{Name: "One Lose Two Win", Code: "one_two", Description: "一輸二贏"},
		entities.PlayType{Name: "Three Way", Code: "three_way", Description: "三路"},
		entities.PlayType{Name: "Odd Even", Code: "odd_even", Description: "單雙"},
		entities.PlayType{Name: "Correct Score", Code: "correct_score", Description: "波膽"},
	}
	for _, item := range playsList {
		fmt.Println("Create Play Type", item.Name, item.Description)
		item = switchPlayType(false, false, item)
		repo.AddPlayTypeByStruct(&item)
		item = switchPlayType(false, true, item)
		repo.AddPlayTypeByStruct(&item)
		item = switchPlayType(true, false, item)
		repo.AddPlayTypeByStruct(&item)
		item = switchPlayType(true, true, item)
		repo.AddPlayTypeByStruct(&item)
	}
}
func switchPlayType(isRunning bool, isParlay bool, pt entities.PlayType) entities.PlayType {
	pt.ID = 0
	pt.IsRunning = isRunning
	pt.IsParlay = isParlay
	return pt
}
