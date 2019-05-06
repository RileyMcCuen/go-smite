package gods

import (
	"encoding/json"
	"io/ioutil"
	G "smite/general"
)

// rawItemPart - unparsed data regarding a specific part of an ability
type rawItemPart struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

// rawItemDescription - unparsed data regarding a specific item (ability, basic attack, etc.)
type rawItemDescription struct {
	Cooldown             string        `json:"cooldown"`
	Cost                 string        `json:"cost"`
	Description          string        `json:"description"`
	Menuitems            []rawItemPart `json:"menuitems"`
	Rankitems            []rawItemPart `json:"rankitems"`
	SecondaryDescription string        `json:"secondaryDescription"`
}

// rawDescription - unparsed data regarding wrapping a rawItemDescription
type rawDescription struct {
	ItemDescription rawItemDescription `json:"itemDescription"`
}

// rawAbility - unparsed data regarding an ability
type rawAbility struct {
	Description rawDescription `json:"Description"`
	ID          int            `json:"Id"`
	Summary     string         `json:"Summary"`
	URL         string         `json:"URL"`
}

// rawGod - Holds data pertaining to unparsed god, intermediate format
type rawGod struct {
	// Ability names
	AbilityName1 string `json:"Ability1"`
	AbilityName2 string `json:"Ability2"`
	AbilityName3 string `json:"Ability3"`
	AbilityName4 string `json:"Ability4"`
	AbilityName5 string `json:"Ability5"`
	// Ability ids
	AbilityID1 int `json:"AbilityId1"`
	AbilityID2 int `json:"AbilityId2"`
	AbilityID3 int `json:"AbilityId3"`
	AbilityID4 int `json:"AbilityId4"`
	AbilityID5 int `json:"AbilityId5"`
	// Ability structs
	AbilityDesciption1 rawAbility `json:"Ability_1"`
	AbilityDesciption2 rawAbility `json:"Ability_2"`
	AbilityDesciption3 rawAbility `json:"Ability_3"`
	AbilityDesciption4 rawAbility `json:"Ability_4"`
	AbilityDesciption5 rawAbility `json:"Ability_5"`
	// God stats
	AttackSpeed                float64 `json:"AttackSpeed"`
	AttackSpeedPerLevel        float64 `json:"AttackSpeedPerLevel"`
	Cons                       string  `json:"Cons"`
	HP5PerLevel                float64 `json:"HP5PerLevel"`
	Health                     int     `json:"Health"`
	HealthPerFive              int     `json:"HealthPerFive"`
	HealthPerLevel             int     `json:"HealthPerLevel"`
	Lore                       string  `json:"Lore"`
	MP5PerLevel                float64 `json:"MP5PerLevel"`
	MagicProtection            int     `json:"MagicProtection"`
	MagicProtectionPerLevel    float64 `json:"MagicProtectionPerLevel"`
	MagicalPower               int     `json:"MagicalPower"`
	MagicalPowerPerLevel       int     `json:"MagicalPowerPerLevel"`
	Mana                       int     `json:"Mana"`
	ManaPerFive                float64 `json:"ManaPerFive"`
	ManaPerLevel               int     `json:"ManaPerLevel"`
	Name                       string  `json:"Name"`
	OnFreeRotation             string  `json:"OnFreeRotation"`
	Pantheon                   string  `json:"Pantheon"`
	PhysicalPower              int     `json:"PhysicalPower"`
	PhysicalPowerPerLevel      int     `json:"PhysicalPowerPerLevel"`
	PhysicalProtection         int     `json:"PhysicalProtection"`
	PhysicalProtectionPerLevel int     `json:"PhysicalProtectionPerLevel"`
	Pros                       string  `json:"Pros"`
	Roles                      string  `json:"Roles"`
	Speed                      int     `json:"Speed"`
	Title                      string  `json:"Title"`
	Type                       string  `json:"Type"`
	// Omitted abilityDescription* fields because they seemed to be duplicates of previous ability fields
	// God basic attack
	BasicAttack rawAbility `json:"basicAttack"`
	// God image urls
	GodAbility1URL string `json:"godAbility1_URL"`
	GodAbility2URL string `json:"godAbility2_URL"`
	GodAbility3URL string `json:"godAbility3_URL"`
	GodAbility4URL string `json:"godAbility4_URL"`
	GodAbility5URL string `json:"godAbility5_URL"`
	GodCardURL     string `json:"godCard_URL"`
	GodIconURL     string `json:"godIcon_URL"`
	// God id
	ID int `json:"id"`
	// Misc data
	LatestGod string `json:"latestGod"`
	RetMsg    string `json:"ret_msg"`
}

func getRawGods() *[]rawGod {
	file, e := ioutil.ReadFile(G.GodsPath)
	G.ErrorHandler(e)
	rgs := make([]rawGod, 0)
	json.Unmarshal(file, &rgs)
	return &rgs
}
