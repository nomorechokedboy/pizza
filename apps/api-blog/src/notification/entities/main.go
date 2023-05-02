package entities

import (
	"api-blog/pkg/entities"
	"time"
)

type NotificationObjectPayload struct {
	NotificationObject
	EntityData string `json:"entityData"`
}

type NotificationObject struct {
	ID                 uint               `json:"id"`
	EntityType         string             `json:"entityType" gorm:"not null"`
	ActionType         string             `json:"actionType" gorm:"not null"`
	EntityID           uint               `json:"entityId" gorm:"not null"`
	Notifications      []Notification     `json:"notifications"`
	NotificationChange NotificationChange `json:"notificationChange"`
	CreatedAt          time.Time          `json:"createdAt" gorm:"not null"`
}

type Notification struct {
	ID                   uint          `json:"id"`
	NotificationObjectID uint          `json:"notificationObjectId" gorm:"not null,index"`
	NotifierID           uint          `json:"notifierId" gorm:"not null,index"`
	Notifier             entities.User `json:"notifier"`
	ReadAt               *time.Time    `json:"readAt"`
}

type NotificationChange struct {
	ID                   uint          `json:"id"`
	NotificationObjectID uint          `json:"notificationObjectId" gorm:"not null,index"`
	ActorID              uint          `json:"actorId" gorm:"not null,index"`
	Actor                entities.User `json:"actor"`
}

type NotificationRequest struct {
	ActionType string
	ActorID    uint
	EntityID   uint
	EntityType string
	EntityData string
	// EntityDataID interface{}
	NotifierID uint
}
