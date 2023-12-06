package dto

type Board struct {
	id      int
	Title   string
	Content string
	Writer  string
}

func (b *Board) GetID() int {
	return b.id
}
