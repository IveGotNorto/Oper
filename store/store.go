package store

type StoreArguments struct {
	Cache   bool
	Verbose bool
	Debug   bool
}

type PasswordStore interface {
	List(string) error
	TreeList(string) error
	Show(string) error
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
	Ascending  = 0
	Descending = 1
)
