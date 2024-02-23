package model

type BoardAdaptor interface {
	SelectBoardAll() ([]*BoardEntity, error)
	SelectBoardByID(id int) (*BoardEntity, error)
	SelectBoardByAuthor(author string) ([]*BoardEntity, error)
	InsertBoard(board *BoardDTO) error
	UpdateBoard(board *BoardDTO) error
	DeleteBoard(id int) error
}

type UserAdaptor interface {
	SelectUserByID(id string) (*UserEntity, error)
	InsertUser(user *UserDTO) error
	UpdateUser(user *UserDTO) error
	DeleteUser(id int) error
}
