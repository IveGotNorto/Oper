package store

type PasswordStore interface {
	List() error
	TreeList() error
	Find([]string) error
	Show(string) error
	Insert(string) error
	Edit() error
	Generate() error
	Remove() error
	Move() error
	Copy() error
	Setup() error
}

type StoreCommands interface {
	GetPassword(string) ([]byte, error)
	GetVaults() ([]byte, error)
	CreatePasswordInContainer() error
	CreatePasswordDefaultContainer() error
	GetPasswordsByContainer(string) ([]byte, error)
}
