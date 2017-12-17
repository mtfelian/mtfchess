package base

// ICells is a board cells
type ICells interface {
	Copy(board IBoard) ICells
}
