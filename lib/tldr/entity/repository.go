package entity

type Repository interface {
	Index() (Index, error)
	Page(name, platform string) (Page, error)
}
