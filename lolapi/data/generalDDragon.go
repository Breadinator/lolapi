package data

type Realm struct {
	N              struct{ Item, Rune, Mastery, Summoner, Champion, ProfileIcon, Map, Language, Sticker string }
	V              string // ? game version?
	L              string // ? locale e.g. "en_GB" for euw
	CDN            string // DataDragon CDN location
	DD             string // ? DataDragon version?
	LG             string
	CSS            string
	ProfileIconMax int
	Store          interface{} // ? what is this?
}
