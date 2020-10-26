package state

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"

	"github.com/EmilyBjartskular/boggle/gamemode/base"
	"github.com/EmilyBjartskular/boggle/network"
	"github.com/EmilyBjartskular/boggle/util/toml"
)

const (
	langsDir string = "configs/languages"
)

func init() {
	// Get ip used for manual connection
	ipAddress = network.GetManualIP()

	// Initialize available languages and their dice configurations
	files, err := ioutil.ReadDir(langsDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		metaPath := fmt.Sprintf("%s/%s/%s", langsDir, file.Name(), "meta.toml")
		metaMap := toml.GetTree(metaPath).ToMap()

		var diceConfigs = make(map[string]base.Dice)
		var boardSizes []base.Size

		lang := metaMap["Meta"].(map[string]interface{})["Language"].(string)

		// Disgusting way of converting the dice configurations
		// from the meta.toml files into the base.Dice type
		diceIList := metaMap["Board"].([]interface{})
		for _, diceConfI := range diceIList {
			var dice base.Dice
			diceConf := diceConfI.(map[string]interface{})["Dice"].([]interface{})
			for _, dieInterface := range diceConf {
				dieI := dieInterface.([]interface{})
				var die base.Die
				for i := 0; i < len(die); i++ {
					die[i] = dieI[i]
				}
				dice = append(dice, die)
			}
			side, err := strconv.Atoi(fmt.Sprintf("%.0f", math.Sqrt(float64(len(dice)))))
			if err != nil {
				panic(err)
			}
			size := fmt.Sprintf("%dx%d", side, side)
			diceConfigs[size] = dice
			boardSizes = append(boardSizes, base.Size{int32(side), int32(side)})
		}

		Languages = append(Languages, &Language{
			Name:        lang,
			DiceConfigs: diceConfigs,
			BoardSizes:  boardSizes,
		})
	}
}
