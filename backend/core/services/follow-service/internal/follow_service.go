package internal

import (
	"services/pkg/common"
)

type FollowService struct {
	Repo *FollowRepository
}

var userServer = "user-service:9090"
var serviceName = "FollowService"

func (svc *FollowService) CreateFollow(req FollowRequest) (*FollowResponse, error) {
	var functionName = "CreateFollow"
	follow := &Follow{
		FollowerID: req.FollowerID,
		FolloweeID: req.FolloweeID,
	}

	if err := svc.Repo.Save(follow); err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during saving follow", err)
	}

	return OriginToResponse(follow)
}

func (svc *FollowService) GetFollow(followID string) (*FollowResponse, error) {
	var functionName = "GetFollow"
	id, err := common.ConvertStringToUint(followID)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	follow, err := svc.Repo.FindByID(id)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during finding follow", err)
	}

	return OriginToResponse(follow)
}

func (svc *FollowService) GetFolloweesByFollowerID(followerID string) ([]*uint, error) {
	var functionName = "GetFolloweesByFollowerID"
	id, err := common.ConvertStringToUint(followerID)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	followeewsID, err := svc.Repo.FindFolloweesByFollowerID(id)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during finding followees by followerID", err)
	}

	return followeewsID, nil
}

func (svc *FollowService) GetFollowersByFolloweeID(followeeID string) ([]*uint, error) {
	var functionName = "GetFollowersByFolloweeID"
	id, err := common.ConvertStringToUint(followeeID)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	followersID, err := svc.Repo.FindFollowersByFolloweeID(id)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during finding followers by followeeID", err)
	}

	return followersID, nil
}

func (svc *FollowService) DeleteFollow(followID string) (bool, error) {
	var functionName = "DeleteFollow"
	id, err := common.ConvertStringToUint(followID)
	if err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	if err := svc.Repo.Delete(id); err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during deleting follow", err)
	}

	return true, nil
}
