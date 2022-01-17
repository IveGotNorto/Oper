package items

import (
	"fmt"
	"os/exec"
	"sort"

	easyjson "github.com/mailru/easyjson"
)

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

func (i *Items) Len() int           { return len(*i) }
func (i *Items) Less(a, b int) bool { return (*i)[a].Overview.Title < (*i)[b].Overview.Title }
func (i *Items) Swap(a, b int)      { (*i)[a], (*i)[b] = (*i)[b], (*i)[a] }

func (i *Items) Retrieve() error {
	out, err := exec.Command("op", "--cache", "list", "items", "--categories", "Login").Output()
	if err != nil {
		return err
	}
	err = i.UnmarshalJSON(out)
	if err != nil {
		return err
	}
	sort.Sort(i)
	return err
}

func (i *Items) Display() {
	for _, item := range *i {
		fmt.Printf("%v\n", item.Overview.Title)
	}
}

func (i *Items) PrettyPrint() error {
	return nil
}

func (i *Items) Find(pass []string) (Item, error) {
	return Item{}, nil
}

func (i *Items) Show(entry string) {
	for j := range *i {
		if entry == (*i)[j].Overview.Title {
			var out []byte
			out, err := exec.Command("op", "--cache", "get", "item", (*i)[j].Uuid, "--fields", "password").Output()
			if err == nil {
				// Simply print the password to the console
				fmt.Printf("%v", string(out))
				return
			}
		}
	}
}

func RetrieveByVault(uuid string) (*MapItems, error) {
	com, err := exec.Command("op", "--cache", "list", "items", "--vault", uuid).Output()
	if err != nil {
		return nil, err
	}
	buf := &Items{}
	err = easyjson.Unmarshal(com, buf)
	if err != nil {
		return nil, err
	}
	tmp := make(MapItems)
	for i := range *buf {
		_, ok := tmp[(*buf)[i].Overview.Title]
		if !ok {
			tmp[(*buf)[i].Overview.Title] = &(*buf)[i]
		}
	}
	return &tmp, err
}
