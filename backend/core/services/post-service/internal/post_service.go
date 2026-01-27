package internal

import (
	"services/pkg/common"
)

type PostService struct {
	Repo *PostRepository
}

var serviceName = "PostService"

func (svc *PostService) CreatePost(req PostRequest) (*PostResponse, error) {
	var functionName = "CreatePost"
	post, postImages, err := PostRequestToOrigins(req)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting request to origin", err)
	}

	// post save
	if err := svc.Repo.postRepo.Save(post); err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during saving post", err)
	}

	// image save
	for _, postImage := range postImages {
		if err := svc.Repo.imageRepo.Save(postImage); err != nil {
			return nil, common.CommitError(serviceName, functionName,
				"error occured during saving post image", err)
		}
	}

	return OriginsToResponse(post, postImages)
}

func (svc *PostService) CreateLike(likeRequest LikeRequest) (bool, error) {
	var functionName = "CreateLike"

	like, err := LikeRequestToOrigin(likeRequest)
	if err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during converting request to origin", err)
	}

	if err := svc.Repo.likeRepo.Save(like); err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during saving like", err)
	}
	return true, nil
}

func (svc *PostService) GetPostById(postID string) (*PostResponse, error) {
	var functionName = "GetPostById"

	id, err := common.ConvertStringToUint(postID)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting postID to uint", err)
	}

	post, err := svc.Repo.postRepo.FindByID(id)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during finding post by postID", err)
	}

	images, err := svc.Repo.FindAllImagesByPostID(id)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during finding post images by postID", err)
	}

	return OriginsToResponse(post, images)
}

func (svc *PostService) UpdatePost(req PostRequest) (*PostResponse, error) {
	var functionName = "UpdatePost"
	post, postImages, err := PostRequestToOrigins(req)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting request to origin", err)
	}

	// post update
	if err := svc.Repo.postRepo.Save(post); err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during updating post", err)
	}

	// image update
	for _, postImage := range postImages {
		if err := svc.Repo.imageRepo.Save(postImage); err != nil {
			return nil, common.CommitError(serviceName, functionName,
				"error occured during updating post image", err)
		}
	}

	return OriginsToResponse(post, postImages)
}

func (svc *PostService) DeletePostById(postID string) (bool, error) {
	var functionName = "DeletePostById"

	id, err := common.ConvertStringToUint(postID)
	if err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	if err := svc.Repo.postRepo.Delete(id); err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during deleting post", err)
	}

	if err := svc.Repo.DeleteAllImagesByPostID(id); err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during deleting post image", err)
	}

	return true, nil
}
