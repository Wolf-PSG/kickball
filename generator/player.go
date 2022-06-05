package generator

import (
	"encoding/json"
	"fmt"
	"go-graphql/main/schema"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"
)

type SkillStruct struct {
	Percentage int
	Affect     string
	Attribute  string
}

type UniqueStruct struct {
	Affect    string
	Attribute string
}

type dict map[string]interface{}

func ReadFile() map[string]interface{} {
	jsonFile, err := os.Open("./generator/skill.json")

	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	result := map[string]interface{}{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

func Skill(modifier string) string {
	jsonResult := ReadFile()
	var skillTypes = []string{
		"buffs",
		"debuffs",
	}
	rand.Seed(time.Now().UnixNano())
	var operator string
	skill := fmt.Sprintf("%s", skillTypes[rand.Intn(len(skillTypes))])
	if skill == "buffs" {
		operator = "+"
	} else {
		operator = "-"
	}
	var skillStruct SkillStruct
	for key, value := range jsonResult {
		if key == skill {
			row := value.(map[string]interface{})
			modifiedRow, exist := row[modifier]
			if exist {
				correction, exist := modifiedRow.(map[string]interface{})
				if exist {
					percentageMap, ok := correction["PERCENTAGES"].([]interface{})
					if ok {
						min, minOk := percentageMap[0].(float64)
						max, maxOk := percentageMap[1].(float64)
						if minOk && maxOk {
							skillStruct.Percentage = GeneratePercentages(int(math.Round(min)), int(math.Round(max)))
						}
					}
					affectingMap, ok := correction["AFFECTING"].([]interface{})
					if ok {
						affectingString, affectingOk := affectingMap[rand.Intn(len(affectingMap))].(string)
						if affectingOk {
							skillStruct.Affect = affectingString
						}
					}
				}
				skillStruct.Attribute = schema.PlayerAttributes[rand.Intn(len(schema.PlayerAttributes))]
			}

		}
	}
	return fmt.Sprintf(`%s%d%% %s %s`, operator, skillStruct.Percentage, skillStruct.Affect, skillStruct.Attribute)
}

func GeneratePercentages(min int, max int) int {
	return rand.Intn((max - min) + min)
}
