package pokeapi

type Pokemon struct {
	Abilities              []Abilities `json:"abilities"`
	BaseExperience         int         `json:"base_experience"`
	Cries                  Cries       `json:"cries"`
	Forms                  []Forms     `json:"forms"`
	Height                 int         `json:"height"`
	HeldItems              []HeldItems `json:"held_items"`
	ID                     int         `json:"id"`
	IsDefault              bool        `json:"is_default"`
	LocationAreaEncounters string      `json:"location_area_encounters"`
	Moves                  []Moves     `json:"moves"`
	Name                   string      `json:"name"`
	Order                  int         `json:"order"`
	PastAbilities          []any       `json:"past_abilities"`
	PastTypes              []any       `json:"past_types"`
	Species                Species     `json:"species"`
	Sprites                Sprites     `json:"sprites"`
	Stats                  []Stats     `json:"stats"`
	Types                  []Types     `json:"types"`
	Weight                 int         `json:"weight"`
}

type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Forms struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type HeldItems struct {
	Item Item `json:"item"`
}
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Moves struct {
	Move Move `json:"move"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type DreamWorld struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}

type Home struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Showdown struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Showdown        `json:"showdown"`
}

type Sprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
	Other            Other  `json:"other"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
