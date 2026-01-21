package internal

type CommentService struct {
	Repo *CommentRepository
}

func (svc *CommentService) CreateComment(req CommentRequest) (*CommentResponse, error) {
	comment := &Comment{
		UserID:  req.UserID,
		PostID:  req.PostID,
		Content: req.Content,
	}

	if err := svc.Repo.Save(comment); err != nil {
		return nil, err
	}

	res := &CommentResponse{
		UserID:    comment.UserID,
		PostID:    comment.PostID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return res, nil
}
