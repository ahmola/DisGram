package internal

import (
	"log/slog"
	"services/pkg/common"
)

type UserService struct {
	Repo *UserRepository
}

var serviceName = "UserService"

func (svc *UserService) CreateUser(req UserRequest) (*UserResponse, error) {
	var functionName = "CreateUser"
	user, err := RequestToOrigin(req)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting request to origin", err)
	}

	if err := svc.Repo.Save(user); err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during saving user", err)
	}

	return OriginToResponse(*user)
}

func (svc *UserService) GetUserByID(userID string) (*UserResponse, error) {
	var functionName = "GetUserByID"
	id, err := common.ConvertStringToUint(userID)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	user, err := svc.Repo.FindByID(id)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"erro occured during finding user by id", err)
	}
	slog.Info("User : ", user.Username)

	return OriginToResponse(*user)
}

func (svc *UserService) GetUserByUsername(username string) (*UserResponse, error) {
	var functionName = "GetUserByUsername"
	user, err := svc.Repo.FindByUsername(username)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during finding user by username", err)
	}

	return OriginToResponse(*user)
}

func (svc *UserService) UpdateUser(req UserRequest) (*UserResponse, error) {
	var functionName = "UpdateUser"
	user, err := RequestToOrigin(req)
	if err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during converting request to origin", err)
	}

	if err := svc.Repo.Update(user); err != nil {
		return nil, common.CommitError(serviceName, functionName,
			"error occured during updating user", err)
	}

	return OriginToResponse(*user)
}

func (svc *UserService) DeleteUser(userID string) (bool, error) {
	var functionName = "DeleteUser"
	id, err := common.ConvertStringToUint(userID)
	if err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during converting string to uint", err)
	}

	if err := svc.Repo.Delete(id); err != nil {
		return false, common.CommitError(serviceName, functionName,
			"error occured during deleting user", err)
	}

	return true, nil
}
