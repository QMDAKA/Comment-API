package comment

type CommentServicer interface {
	GetAll()
}

type Comment struct {
}

func NewComment() *Comment {
	return &Comment{}
}

func (c *Comment) GetAll() {

}
