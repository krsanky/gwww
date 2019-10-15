package phrase

import "fmt"

type Phrase struct {
	Id     int
	Phrase string
	Path   string
}

func (p *Phrase) String() string {
	return fmt.Sprintf("<Phrase id:%d %s...>", p.Id, p.Phrase[:13])
}

