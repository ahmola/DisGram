package internal

import (
	"services/pkg/common"
)

type CommentService struct {
	Repo *CommentRepository
}

var serviceName = "CommentService"

func (svc *CommentService) CreateComment(req CommentRequest) (*CommentResponse, error) {
	var functionName = "CreateComment"
	comment, err := RequestToOrigin(req)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting request to origin", err)
	}

	if err := svc.Repo.Save(comment); err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during saving comment", err)
	}

	return OriginToResponse(*comment)
}

func (svc *CommentService) GetCommentsByPostID(postID string) ([]*CommentResponse, error) {
	var functionName = "GetCommentsByPostID"
	id, err := ConvertStringToUint(postID)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting postID to uint: %w", err)
	}

	comments, err := svc.Repo.FindByPostID(id)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during finding comments by postID", err)
	}

	res := make([]*CommentResponse, len(comments))
	for i, comment := range comments {
		res[i], err = OriginToResponse(comment)
		if err != nil {
			return nil, common.CommitError(serviceName, functionName,
				"error occured during making response ", err)
		}
	}

	return res, nil
}

func (svc *CommentService) GetCommentByID(commentID string) (*CommentResponse, error) {
	var functionName = "GetCommentById"
	id, err := ConvertStringToUint(commentID)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	comment, err := svc.Repo.FindByID(id)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	return OriginToResponse(*comment)
}

func (svc *CommentService) UpdateComment(req CommentRequest) (*CommentResponse, error) {
	var functionName = "UpdateComment"
	comment, err := RequestToOrigin(req)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting request to origin", err)
	}

	if err := svc.Repo.Update(comment); err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during updating comment", err)
	}

	return OriginToResponse(*comment)
}

func (svc *CommentService) DeleteComment(commentID string) (bool, error) {
	var functionName = "DeleteComment"
	id, err := ConvertStringToUint(commentID)
	if err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	if err := svc.Repo.Delete(id); err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during deleting comment", err)
	}

	return true, nil
}
