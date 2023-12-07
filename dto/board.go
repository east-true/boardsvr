package dto

type Board struct {
	id      int
	Title   string
	Content string
	Writer  string // TODO : edit author(user dto array)
}

func (b *Board) GetID() int {
	return b.id
}
