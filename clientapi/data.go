package clientapi

// Various strucs for the data

type DataEntry struct {
	ID                                          interface{} // string or int
	DisplayName, RawDescription, RawDisplayName string
}

type Ability struct {
	DataEntry
	AbilityLevel int
}
type Abilities struct{ Passive, Q, W, E, R Ability }
type Runes struct {
	GeneralRunes, StatRunes                      []DataEntry
	Keystone, PrimaryRuneTree, SecondaryRuneTree DataEntry
}

type Event struct {
	EventID   int
	EventName string
	EventTime float64
}

type ActivePlayer struct {
	Abilities          Abilities
	ChampionStats      map[string]interface{} // floats or strings
	CurrentGold        float64
	FullRunes          Runes
	Level              int
	SummonerName       string
	TeamRelativeColors bool
}

type Scores struct {
	Assists, CreepScore, Deaths, Kills int
	WardScore                          float64
}

type Player struct {
	ChampionName    string
	IsBot           bool
	IsDead          bool
	Items           []Item
	Level           int
	Position        string
	RawChampionName string
	RespawnTimer    float32
	Runes           Runes
	Scores          Scores
	SkinID          int
	SummonerName    string
	SummonerSpells  struct{ SummonerSpellOne, SummonerSpellTwo DataEntry }
	Team            string
}

type AllGameData struct {
	ActivePlayer ActivePlayer
	AllPlayers   []Player
	Events       Events
	GameData     GameStats
}

type Events struct {
	Events []Event
}

type GameStats struct {
	GameMode, MapName, MapTerrain string
	GameTime                      float64
	MapNumber                     int
}

type Item struct {
	CanUse, Consumable                          bool
	Count, ItemID, Price, Slot                  int
	DisplayName, RawDescription, RawDisplayName string
}
