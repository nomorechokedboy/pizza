package entities

import (
	"api-blog/pkg/entities"
	"time"
)

type NotificationObjectReal struct {
	ID                 uint
	EntityData         string
	ActionType         string
	Notifications      []Notification
	NotificationChange NotificationChange
	CreatedAt          time.Time
}

type NotificationObjectDB struct {
	ID                 uint
	EntityData         string
	ActionType         string
	Notification       Notification
	NotificationChange NotificationChange
	CreatedAt          time.Time
}

func (n *NotificationObjectDB) toNotificationObject() *NotificationObjectReal {
	return &NotificationObjectReal{ID: n.ID,
		CreatedAt:          n.CreatedAt,
		EntityData:         n.EntityData,
		ActionType:         n.ActionType,
		Notifications:      []Notification{n.Notification},
		NotificationChange: n.NotificationChange,
	}
}

func (n *NotificationObjectReal) groupByField() uint {
	return n.ID
}

func (n *NotificationObjectReal) vecField() []Notification {
	return n.Notifications
}

func (n *NotificationObjectReal) groupBy(items []NotificationObjectDB) (m map[uint]NotificationObjectReal) {
	for _, db_item := range items {
		if o, ok := m[db_item.ID]; ok {
			o.Notifications = append(o.Notifications, db_item.Notification)
		} else {
			m[n.groupByField()] = *db_item.toNotificationObject()
		}
	}

	return m
}

func (n *NotificationObjectReal) flatten(items []NotificationObjectDB) (res []NotificationObjectReal) {
	for _, el := range n.groupBy(items) {
		res = append(res, el)
	}

	return res
}

func (n *NotificationObjectReal) flattenOne(items []NotificationObjectDB) *NotificationObjectReal {
	res := []NotificationObjectReal{}
	for _, el := range n.groupBy(items) {
		res = append(res, el)
	}

	return &res[0]
}

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
