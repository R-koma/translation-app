package services

import (
	"github.com/R-koma/translation-app/backend/models"
	"github.com/R-koma/translation-app/backend/repositories"
)

type IFriendRequestService interface {
	CreateFriendRequest(senderID uint, receiverID uint) error
	GetFriendRequestsByReceiverID(receiverID uint) ([]models.FriendRequest, error)
	UpdateFriendRequestStatus(id uint, status string) error
	GetMyFriends(userID uint) ([]models.User, error)
}

type friendRequestService struct {
	repository repositories.IFriendRequestRepository
}

func NewFriendRequestService(repository repositories.IFriendRequestRepository) IFriendRequestService {
	return &friendRequestService{repository: repository}
}

func (s *friendRequestService) CreateFriendRequest(senderID uint, receiverID uint) error {
	req := models.FriendRequest{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Status:     "pending",
	}
	return s.repository.CreateRequest(&req)
}

func (s *friendRequestService) GetFriendRequestsByReceiverID(receiverID uint) ([]models.FriendRequest, error) {
	return s.repository.FindByReceiverID(receiverID)
}

func (s *friendRequestService) UpdateFriendRequestStatus(id uint, status string) error {
	if status != "accepted" && status != "rejected" {
		return nil
	}
	return s.repository.UpdateStatus(id, status)
}

func (s *friendRequestService) GetMyFriends(userID uint) ([]models.User, error) {
	return s.repository.FindAcceptedFriendUsers(userID)
}
