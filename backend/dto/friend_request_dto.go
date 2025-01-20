package dto

type CreateFriendRequestDto struct {
	ReceiverID uint `json:"receiver_id" binding:"required"`
}

type UpdateFriendRequestStatusDto struct {
	Status string `json:"status" binding:"required,oneof=accepted rejected"`
}
