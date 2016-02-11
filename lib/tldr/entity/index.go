package entity

type Index interface {
	PlatformsFor(name string) []string
	Commands() map[string][]string
}

func NewIndex(commands map[string][]string) Index {
	return &index{
		commands: commands,
	}
}

type index struct {
	commands map[string][]string
}

func (i *index) PlatformsFor(name string) []string {
	return i.commands[name]
}

func (i *index) Commands() map[string][]string {
	return i.commands
}
