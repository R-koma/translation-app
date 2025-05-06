package repositories

import (
	"github.com/R-koma/translation-app/backend/models"
	"gorm.io/gorm"
)

type IFriendRequestRepository interface {
	CreateRequest(req *models.FriendRequest) error
	FindByReceiverID(receiverID uint) ([]models.FriendRequest, error)
	UpdateStatus(id uint, status string) error
	FindAcceptedFriendUsers(userID uint) ([]models.User, error)
}

type friendRequestRepository struct {
	db *gorm.DB
}

func NewFriendRequestRepository(db *gorm.DB) IFriendRequestRepository {
	return &friendRequestRepository{db: db}
}

func (r *friendRequestRepository) CreateRequest(req *models.FriendRequest) error {
	result := r.db.Create(req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *friendRequestRepository) FindByReceiverID(receiverID uint) ([]models.FriendRequest, error) {
	var requests []models.FriendRequest
	result := r.db.Find(&requests, "receiver_id = ?", receiverID)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}

func (r *friendRequestRepository) UpdateStatus(id uint, status string) error {
	result := r.db.Model(&models.FriendRequest{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *friendRequestRepository) FindAcceptedFriendUsers(userID uint) ([]models.User, error) {
	var users []models.User
	result := r.db.
		Model(&models.User{}).
		Distinct().
		Select("users.*").
		Joins("INNER JOIN friend_requests ON (users.id = friend_requests.sender_id AND friend_requests.receiver_id = ?) OR (users.id = friend_requests.receiver_id AND friend_requests.sender_id = ?)", userID, userID).
		Where("friend_requests.status = ?", "accepted").
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
