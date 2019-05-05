package items

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	G "smite/general"
	"strconv"
)

// aflagStringToBool - convert aflag to bool value
func aflagStringToBool(s string) bool {
	if s == "y" {
		return true
	}
	return false
}

// MenuItemDescription - menu item description represented as int
type MenuItemDescription int

// All currently existing MenuItemDescriptions
const (
	NoMenuItemDescription MenuItemDescription = -1
	AttackSpeed           MenuItemDescription = 0
	MagicalLifesteal      MenuItemDescription = 1
	MP5                   MenuItemDescription = 2
	MagicalProtection     MenuItemDescription = 3
	MovementSpeed         MenuItemDescription = 4
	Health                MenuItemDescription = 5
	PhysicalProtection    MenuItemDescription = 6
	MagicalPower          MenuItemDescription = 8
	CrowdControlReduction MenuItemDescription = 9
	PhysicalPenetration   MenuItemDescription = 10
	CriticalStrikeChance  MenuItemDescription = 11
	PhysicalPower         MenuItemDescription = 12
	Penetration           MenuItemDescription = 13
	CooldownReduction     MenuItemDescription = 14
	MagicalPenetration    MenuItemDescription = 15
	HP5                   MenuItemDescription = 16
	PhysicalLifesteal     MenuItemDescription = 17
	Mana                  MenuItemDescription = 18
)

// menuItemTextToMenuItemDescription - converts menu item text to MenuItemDescription
func menuItemTextToMenuItemDescription(s string) MenuItemDescription {
	switch s {
	case "Critical Strike Chance":
		return CriticalStrikeChance
	case "Magical Lifesteal":
		return MagicalLifesteal
	case "MP5":
		return MP5
	case "Magical Power", "Magical power", "Magical Power ":
		return MagicalPower
	case "Crowd Control Reduction", "Crowd Control Reduction:":
		return CrowdControlReduction
	case "Physical Protection":
		return PhysicalProtection
	case "Physical Penetration":
		return PhysicalPenetration
	case "Magical Protection":
		return MagicalProtection
	case "Penetration":
		return Penetration
	case "Attack Speed":
		return AttackSpeed
	case "Physical Power", "Physical power":
		return PhysicalPower
	case "Physical Lifesteal":
		return PhysicalLifesteal
	case "Mana":
		return Mana
	case "Cooldown Reduction":
		return CooldownReduction
	case "Health":
		return Health
	case "Movement Speed":
		return MovementSpeed
	case "Magical Penetration":
		return MagicalPenetration
	case "HP5":
		return HP5
	default:
		return NoMenuItemDescription
	}
}

// ValueModifier - the way to modify the value of the menu item
type ValueModifier int

// Ways to modify menu items, some items do not have modification,
// while others modify by constant or percentage of currenct value
const (
	NoValueModifier ValueModifier = 0
	Constant        ValueModifier = 1
	Percent         ValueModifier = 2
)

// valueTextToValueAndValueModifier - converts value text to Value and ValueModifier
func valueTextToValueAndValueModifier(s string) (v int, vm ValueModifier) {
	r, _ := regexp.Compile("[\\+]")
	s = r.ReplaceAllString(s, "")
	r, _ = regexp.Compile("[\\%]")
	if r.MatchString(s) {
		vm = Percent
		s = r.ReplaceAllString(s, "")
	} else {
		vm = Constant
	}
	v, e := strconv.Atoi(s)
	if e != nil {
		vm = NoValueModifier
		v = 0
	}
	return v, vm
}

// rawMenuItemToMenuItem - converts rawMenuItemToMenuItem to menuItem
func rawMenuItemToMenuItem(rmi rawMenuItem) menuItem {
	var mi menuItem
	mi.Description = menuItemTextToMenuItemDescription(rmi.Description)
	v, vm := valueTextToValueAndValueModifier(rmi.Value)
	mi.Value = v
	mi.Modifier = vm
	return mi
}

// allRawMenuItemToMenuItemSlice - converts slice of rawMenuItem to slice of menuItem
func allRawMenuItemsToMenuItemSlice(rmis []rawMenuItem) []menuItem {
	ret := make([]menuItem, len(rmis))
	for i, rmi := range rmis {
		ret[i] = rawMenuItemToMenuItem(rmi)
	}
	return ret
}

// menuItem - represents a menu item description and value within an item's description
type menuItem struct {
	Description MenuItemDescription
	Modifier    ValueModifier
	Value       int
}

// rawItemDescriptionToItemDescription - converts rawItemDescription to itemDescription
func rawItemDescriptionToItemDescription(rid rawItemDescription) itemDescription {
	var id itemDescription
	id.Description = rid.Description
	id.MenuItems = allRawMenuItemsToMenuItemSlice(rid.MenuItems)
	id.SecondaryDescription = rid.SecondaryDescription
	return id
}

// itemDescription - represents an item description of a whole item
type itemDescription struct {
	Description          string
	MenuItems            []menuItem
	SecondaryDescription string
}

// Role - role of god represented by int
type Role int

// All of the roles of Smite gods
const (
	NoRole   Role = -1
	Assassin Role = 0
	Hunter   Role = 1
	Mage     Role = 2
	Warrior  Role = 3
	Guardian Role = 4
)

// restrictedRoleTextToRoleSlice - converts restricted role text to slice of Role
func restrictedRoleTextToRoleSlice(s string) []Role {
	r, _ := regexp.Compile("([a-zA-Z]+)")
	o := r.FindAllString(s, -1)
	if o == nil {
		return nil
	}
	ret := make([]Role, 0)
	for _, os := range o {
		switch os {
		case "assassin":
			ret = append(ret, Assassin)
		case "hunter":
			ret = append(ret, Hunter)
		case "mage":
			ret = append(ret, Mage)
		case "warrior":
			ret = append(ret, Warrior)
		case "guardian":
			ret = append(ret, Guardian)
		default:
			ret = append(ret, NoRole)
		}
	}
	allNegative := true
	for _, os := range ret {
		if os != NoRole {
			allNegative = false
		}
	}
	if allNegative {
		return []Role{NoRole}
	}
	return ret
}

// Type - type of item that an item is represented as int
type Type int

// All current item types
const (
	NoType     Type = -1
	Active     Type = 0
	Consumable Type = 1
	Item       Type = 2
)

// typeTextToType - converts type text to Type
func typeTextToType(s string) Type {
	switch s {
	case "Item":
		return Item
	case "Consumable":
		return Consumable
	case "Active":
		return Active
	default:
		return NoType
	}
}

// item - Holds data pertaining to parsed item
type item struct {
	ActiveFlag      bool
	ChildItemID     int
	DeviceName      string
	IconID          int
	ItemDescription itemDescription
	ItemID          int
	ItemTier        int
	Price           int
	RestrictedRoles []Role
	RootItemID      int
	ShortDesc       string
	StartingItem    bool
	Type            Type
	ItemIconURL     string
	RetMessage      string
}

// rawMenuItem - represents a menu item description and value within an item's description
type rawMenuItem struct {
	Description string `json:"Description"`
	Value       string `json:"Value"`
}

// rawItemDescription - represents an item description of a whole item
type rawItemDescription struct {
	Description          string        `json:"Description"`
	MenuItems            []rawMenuItem `json:"MenuItems"`
	SecondaryDescription string        `json:"SecondaryDescription"`
}

// rawItem - Holds data pertaining to unparsed item, intermediate format
type rawItem struct {
	ActiveFlag      string             `json:"ActiveFlag"`
	ChildItemID     int                `json:"ChildItemId"`
	DeviceName      string             `json:"DeviceName"`
	IconID          int                `json:"IconId"`
	ItemDescription rawItemDescription `json:"ItemDescription"`
	ItemID          int                `json:"ItemId"`
	ItemTier        int                `json:"ItemTier"`
	Price           int                `json:"Price"`
	RestrictedRoles string             `json:"RestrictedRoles"`
	RootItemID      int                `json:"RootItemId"`
	ShortDesc       string             `json:"ShortDesc"`
	StartingItem    bool               `json:"StartingItem"`
	Type            string             `json:"Type"`
	ItemIconURL     string             `json:"itemIcon_URL"`
	RetMessage      string             `json:"ret_msg"`
}

// rawItemToItem - converts a rawItem to an item
func rawItemToItem(ri rawItem) item {
	var i item
	i.ActiveFlag = aflagStringToBool(ri.ActiveFlag)
	i.ChildItemID = ri.ChildItemID
	i.DeviceName = ri.DeviceName
	i.IconID = ri.IconID
	i.ItemDescription = rawItemDescriptionToItemDescription(ri.ItemDescription)
	i.ItemID = ri.ItemID
	i.ItemTier = ri.ItemTier
	i.Price = ri.Price
	i.RestrictedRoles = restrictedRoleTextToRoleSlice(ri.RestrictedRoles)
	i.RootItemID = ri.RootItemID
	i.ShortDesc = ri.ShortDesc
	i.StartingItem = ri.StartingItem
	i.Type = typeTextToType(ri.Type)
	i.ItemIconURL = ri.ItemIconURL
	i.RetMessage = ri.RetMessage
	return i
}

// allRawItemsToItemSlice - converts slice of rawItem to slice of item
func allRawItemsToItemSlice(ris []rawItem) []item {
	is := make([]item, len(ris))
	for i, ri := range ris {
		is[i] = rawItemToItem(ri)
	}
	return is
}

func getRawItems() []rawItem {
	file, e := ioutil.ReadFile(G.ItemsPath)
	G.ErrorHandler(e)
	ris := make([]rawItem, 0)
	json.Unmarshal(file, &ris)
	return ris
}
