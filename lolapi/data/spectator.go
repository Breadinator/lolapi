package data

// List of featured games gotten from the endpoint GetGamesFeatured
type FeaturedGames struct {
	GameList              []FeaturedGameInfo // The list of featured games
	ClientRefreshInterval int64              // The suggested interval to wait before requesting FeaturedGames again
}

// Info on a given game listed by featured games
type FeaturedGameInfo struct {
	GameMode          string                // The game mode (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO)
	GameLength        int64                 // The amount of time in seconds that has passed since the game started
	MapID             int64                 // The ID of the map
	GameType          string                // The game type (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME)
	BannedChampions   []BannedChampion      // Banned champion information
	GameID            int64                 // The ID of the game
	Observer          Observer              // The observer information
	GameQueueConfigID int64                 // The queue type
	GameStartTime     int64                 // The game start time represented in epoch milliseconds
	Participants      []FeaturedParticipant // The participant information
	PlatformID        string                // The ID of the platform on which the game is being played
}

// The observer information
type Observer struct {
	EncryptionKey string // Key used to decrypt the spectator grid game data for playback
}

// Banned champion information
type BannedChampion struct {
	PickTurn   int64 // The turn during which the champion was banned
	ChampionID int64 // The ID of the banned champion
	TeamID     int64 // The ID of the team that banned the champion
}

// The participant information of a featured game
type FeaturedParticipant struct {
	Bot           bool   // Flag indicating whether or not this participant is a bot
	Spell1ID      int64  // The ID of the first summoner spell used by this participant
	Spell2ID      int64  // The ID of the second summoner spell used by this participant
	ProfileIconID int64  // The ID of the profile icon used by this participant
	SummonerName  string // The summoner name of this participant
	ChampionID    int64  // The ID of the champion played by this participant
	TeamID        int64  // The ID of this participant, indicating the participant's team
}

type CurrentGameInfo struct {
	GameID            int64                    // The ID of the game
	GameType          string                   // The game type
	GameStartTime     int64                    // The game start time represented in epoch milliseconds
	MapID             int64                    // The ID of the map
	GameLength        int64                    // The amount of time in seconds that has passed since the game started
	PlatformID        string                   // The ID of the platform on which the game is being played
	GameMode          string                   // The game mode
	BannedChampions   []BannedChampion         // Banned champion information
	GameQueueConfigID int64                    // The queue type (queue types are documented on the Game Constants page)
	Observers         Observer                 // The observer information
	Participants      []CurrentGameParticipant // The participant information
}

// The participant information
type CurrentGameParticipant struct {
	ChampionID               int64                     // The ID of the champion played by this participant
	Perks                    Perks                     // Perks/Runes Reforged Information
	ProfileIconID            int64                     // The ID of the profile icon used by this participant
	Bot                      bool                      // Flag indicating whether or not this participant is a bot
	TeamID                   int64                     // The team ID of this participant, indicating the participant's team
	SummonerName             string                    // The summoner name of this participant
	SummonerID               string                    // The encrypted summoner ID of this participant
	Spell1ID                 int64                     // The ID of the first summoner spell used by this participant
	Spell2ID                 int64                     // The ID of the second summoner spell used by this participant
	GameCustomizationObjects []GameCustomizationObject // List of Game Customizations
}

// Perks/Runes Reforged Information
type Perks struct {
	PerkIDs       []int64 // IDs of the perks/runes assigned
	PerkStyle     int64   // Primary runes path
	PerksSubStype int64   // Secondary runes path
}

// List of Game Customizations
type GameCustomizationObject struct {
	Category string // Category identifier for the Game Customization
	Content  string // Game customization content
}
