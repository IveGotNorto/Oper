package store

type StoreArguments struct {
	Cache   bool
	Verbose bool
	Debug   bool
}

type PasswordStore interface {
	List(string) error
	TreeList(string) error
	Find([]string) error
	Show(string) error
	Insert(string) error
	Edit() error
	Generate() error
	Remove() error
	Move() error
	Copy() error
	Setup(StoreArguments) error
}

type StoreCommands interface {
	GetPassword(string) ([]byte, error)
	GetVaults() ([]byte, error)
	CreatePasswordInContainer() error
	CreatePasswordDefaultContainer() error
	GetPasswordsByContainer(string) ([]byte, error)
}

const (
	Ascending int = 0
	Descending
)
