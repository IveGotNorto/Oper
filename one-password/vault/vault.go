package vault

type ItemMap interface {
}

//easyjson:json
type Vault struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Items *ItemMap
	Length int
}

func (v *Vault) GetItems() {


}

func (v *Vault) List () {

}

func (v *Vault) Add () {

}

func (v *Vault) Edit () {

}

func (v *Vault) Delete () {

}

// Possibly a create?
