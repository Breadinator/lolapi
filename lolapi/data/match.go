package data

type MatchDTO struct {
	Metadata MetadataDTO // Match metadata.
	Info     InfoDTO     // Match info.
}

type MetadataDTO struct {
	DataVersion  string   // Match data version.
	MatchID      string   // Match ID.
	Participants []string // A list of participant PUUIDs
}

type InfoDTO struct {
	GameCreation       int64            // Unix timestamp for when the game is created on the game server (i.e., the loading screen).
	GameDuration       int64            // Prior to patch 11.20, this field returns the game length in milliseconds calculated from gameEndTimestamp - gameStartTimestamp. Post patch 11.20, this field returns the max timePlayed of any participant in the game in seconds, which makes the behavior of this field consistent with that of match-v4. The best way to handling the change in this field is to treat the value as milliseconds if the gameEndTimestamp field isn't in the response and to treat the value as seconds if gameEndTimestamp is in the response.
	GameEndTimestamp   int64            // Unix timestamp for when match ends on the game server. This timestamp can occasionally be significantly longer than when the match "ends". The most reliable way of determining the timestamp for the end of the match would be to add the max time played of any participant to the gameStartTimestamp. This field was added to match-v5 in patch 11.20 on Oct 5th, 2021.
	GameID             int64            //
	GameMode           string           // Refer to the Game Constants documentation.
	GameName           string           //
	GameStartTimestamp int64            // Unix timestamp for when match starts on the game server.
	GameType           string           //
	GameVersion        string           // The first two parts can be used to determine the patch a game was played on.
	MapID              int              // Refer to the Game Constants documentation.
	Participants       []ParticipantDTO //
	PlatformID         string           // Platform where the match was played.
	QueueID            int              // Refer to the Game Constants documentation.
	Teams              []TeamDTO        //
	TournamentCode     string           // Tournament code used to generate the match. This field was added to match-v5 in patch 11.13 on June 23rd, 2021.
}

type ParticipantDTO struct {
	Assists                        int
	BaronKills                     int
	BountyLevel                    int
	ChampExperience                int
	ChampLevel                     int
	ChampionID                     int // Prior to patch 11.4, on Feb 18th, 2021, this field returned invalid championIds. We recommend determining the champion based on the championName field for matches played prior to patch 11.4.
	ChampionName                   string
	ChampionTransform              int // This field is currently only utilized for Kayn's transformations. (Legal values: 0 - None, 1 - Slayer, 2 - Assassin)
	ConsumablesPurchased           int
	DamageDealtToBuildings         int
	DamageDealtToObjectives        int
	DamageDealtToTurrets           int
	DamageSelfMitigated            int
	Deaths                         int
	DetectorWardsPlaced            int
	DoubleKills                    int
	DragonKills                    int
	FirstBloodAssist               bool
	FirstBloodKill                 bool
	FirstTowerAssist               bool
	FirstTowerKill                 bool
	GameEndedInEarlySurrender      bool
	GameEndedInSurrender           bool
	GoldEarned                     int
	GoldSpent                      int
	IndividualPosition             string //Both individualPosition and teamPosition are computed by the game server and are different versions of the most likely position played by a player. The individualPosition is the best guess for which position the player actually played in isolation of anything else. The teamPosition is the best guess for which position the player actually played if we add the constraint that each team must have one top player, one jungle, one middle, etc. Generally the recommendation is to use the teamPosition field over the individualPosition field.
	InhibitorKills                 int
	InhibitorTakedowns             int
	InhibitorsLost                 int
	Item0                          int
	Item1                          int
	Item2                          int
	Item3                          int
	Item4                          int
	Item5                          int
	Item6                          int
	ItemsPurchased                 int
	KillingSprees                  int
	Kills                          int
	Lane                           string
	LargestCriticalStrike          int
	LargestKillingSpree            int
	LargestMultiKill               int
	LongestTimeSpentLiving         int
	MagicDamageDealt               int
	MagicDamageDealtToChampions    int
	MagicDamageTaken               int
	NeutralMinionsKilled           int
	NexusKills                     int
	NexusTakedowns                 int
	NexusLost                      int
	ObjectivesStolen               int
	ObjectivesStolenAssists        int
	ParticipantID                  int
	PentaKills                     int
	Perks                          PerksDTO
	PhysicalDamageDealt            int
	PhysicalDamageDealtToChampions int
	PhysicalDamageTaken            int
	ProfileIcon                    int
	PUUID                          string
	QuadraKills                    int
	RiotIDName                     string
	RiotIDTagline                  string
	Role                           string
	SightWardsBoughtInGame         int
	Spell1Casts                    int
	Spell2Casts                    int
	Spell3Casts                    int
	Spell4Casts                    int
	Summoner1Casts                 int
	Summoner1ID                    int
	Summoner2Casts                 int
	Summoner2ID                    int
	SummonerID                     string
	SummonerLevel                  int
	SummonerName                   string
	TeamEarlySurrendered           bool
	TeamID                         int
	TeamPosition                   string // Both individualPosition and teamPosition are computed by the game server and are different versions of the most likely position played by a player. The individualPosition is the best guess for which position the player actually played in isolation of anything else. The teamPosition is the best guess for which position the player actually played if we add the constraint that each team must have one top player, one jungle, one middle, etc. Generally the recommendation is to use the teamPosition field over the individualPosition field.
	TimeCCingOthers                int
	TimePlayed                     int
	TotalDamageDealt               int
	TotalDamageDealtToChampions    int
	TotalDamageShieldedOnTeammates int
	TotalDamageTaken               int
	TotalHeal                      int
	TotalHealsOnTeammates          int
	TotalMinionsKilled             int
	TotalTimeCCDealt               int
	TotalTimeSpentDead             int
	TotalUnitsHealed               int
	TripleKills                    int
	TrueDamageDealt                int
	TrueDamageDealtToChampions     int
	TrueDamageTaken                int
	TurretKills                    int
	TurretTakedowns                int
	TurretsLost                    int
	UnrealKills                    int
	VisionScore                    int
	VisionWardsBoughtInGame        int
	WardsKilled                    int
	WardsPlaced                    int
	Win                            bool
}

type PerksDTO struct {
	StatPerks struct{ Defense, Flex, Offense int }
	Styles    struct {
		Description string
		Selections  []struct{ Perk, Var1, Var2, Var3 int }
		Style       int
	}
}

type TeamDTO struct {
	Bans       []struct{ ChampionID, PickTurn int }
	Objectives struct{ Baron, Champion, Dragon, Inhibitor, RiftHrald, Tower ObjectiveDTO }
	TeamID     int
	Win        bool
}

type ObjectiveDTO struct {
	First bool
	Kills int
}
