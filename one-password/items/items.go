package items

//easyjson:json
type URL struct {
	L string `json:"l"`
	U string `json:"u"`
}

//easyjson:json
type Overview struct {
	URLs  []URL    `json:"URLs"`
	Ainfo string   `json:"ainfo"`
	Pbe   float64  `json:"pbe"`
	Pgrng bool     `json:"pgrng"`
	Ps    int      `json:"ps"`
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
	Url   string   `json:"url"`
}

//easyjson:json
type Item struct {
	Uuid         string   `json:"uuid"`
	TemplateUuid string   `json:"templateUuid"`
	Trashed      string   `json:"trashed"`
	CreatedAt    string   `json:"createdAt"`
	UpdatedAt    string   `json:"updatedAt"`
	ChangerUuid  string   `json:"changerUuid"`
	ItemVersion  int      `json:"itemVersion"`
	VaultUuid    string   `json:"vaultUuid"`
	Overview     Overview `json:"overview"`
}

//easyjson:json
type Items []Item

type MapItems map[string]*Item

func RetrieveByVault(uuid string) (*MapItems, error) {
	buf, err := getItems(uuid)
	tmp := make(MapItems)
	for i := range *buf {
		_, ok := tmp[(*buf)[i].Overview.Title]
		if !ok {
			tmp[(*buf)[i].Overview.Title] = &(*buf)[i]
		}
	}
	return &tmp, err
}
