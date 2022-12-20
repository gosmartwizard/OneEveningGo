package server

import "sync"

type UserPayload struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type TweetId struct {
	Id int `json:"ID"`
}

type UserPayloadRepository interface {
	AddUserPayload(UserPayload) (int, error)
	GetUserPayloads() ([]UserPayload, error)
}

type UserInMemoryRepository struct {
	lock        sync.RWMutex
	userPayload []UserPayload
}

func (umr *UserInMemoryRepository) AddUserPayload(up UserPayload) (int, error) {
	umr.lock.RLock()
	defer umr.lock.RUnlock()

	umr.userPayload = append(umr.userPayload, up)
	return len(umr.userPayload), nil
}

func (umr *UserInMemoryRepository) GetUserPayloads() ([]UserPayload, error) {
	umr.lock.RLock()
	defer umr.lock.RUnlock()
	return umr.userPayload, nil
}

type UserSecondaryMemoryRepository struct {
	userPayload []UserPayload
}

func (umr *UserSecondaryMemoryRepository) AddUserPayload(up UserPayload) (int, error) {
	umr.userPayload = append(umr.userPayload, up)
	return len(umr.userPayload), nil
}
