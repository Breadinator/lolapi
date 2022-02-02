package data

// Represents a summoner
type SummonerDTO struct {
	AccountID     string // Encrypted account ID. Max length 56 characters.
	Name          string // Summoner name.
	ID            string // Encrypted summoner ID. Max length 63 characters.
	PUUID         string // Encrypted PUUID. Exact length of 78 characters.
	ProfileIconID int64  // ID of the summoner icon associated with the summoner.
	SummonerLevel int64  // Summoner level associated with the summoner.
	RevisionDate  int64  // Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: summoner name change, summoner level change, or profile icon change.
}
