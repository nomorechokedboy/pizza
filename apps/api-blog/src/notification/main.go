package notification

import (
	"api-blog/pkg/entities"
	notificationEntities "api-blog/src/notification/entities"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NotifyRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewNotifyRepository(db *gorm.DB, rdb *redis.Client) *NotifyRepository {
	return &NotifyRepository{db: db, rdb: rdb}
}

func (n *NotifyRepository) Notify(req notificationEntities.NotificationRequest) {
	notificationObject := notificationEntities.NotificationObject{
		EntityID:   req.EntityID,
		EntityType: req.EntityType,
		ActionType: req.ActionType,
	}
	if err := n.db.
		Create(&notificationObject).Error; err != nil {
		log.Println("Error happen sometimes lol", err)
	}

	notification := notificationEntities.Notification{
		NotifierID:           req.NotifierID,
		NotificationObjectID: notificationObject.ID,
		Notifier:             entities.User{Id: req.NotifierID},
	}

	notificationChange := notificationEntities.NotificationChange{
		NotificationObjectID: notificationObject.ID,
		ActorID:              req.ActorID,
	}
	if err := n.db.
		Transaction(func(tx *gorm.DB) error {
			if err := tx.
				Clauses(clause.Returning{}).
				Create(&notification).Error; err != nil {
				return err
			}
			if err := tx.
				Clauses(clause.Returning{}).
				Create(&notificationChange).Error; err != nil {
				return err
			}

			return nil
		}); err != nil {
		log.Println(err)
	}

	notificationObject.Notifications = []notificationEntities.Notification{notification}
	notificationObject.NotificationChange = notificationChange
	notificationPayload := notificationEntities.NotificationObjectPayload{
		NotificationObject: notificationObject,
		EntityData:         req.EntityData,
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	payload, err := json.
		Marshal(notificationPayload)
	if err != nil {
		log.Println(err)
	}
	if err := n.rdb.
		Publish(ctx, "notification", payload).Err(); err != nil {
		log.Println("Help me", err)
	}
}

func (n *NotifyRepository) DeleteNotification(req notificationEntities.NotificationRequest) {
	notificationObject := notificationEntities.NotificationObject{
		EntityID:   req.EntityID,
		EntityType: req.EntityType,
		ActionType: req.ActionType,
	}
	if err := n.db.
		Where("action_type = ?", req.ActionType).
		Where("entity_type = ?", req.EntityType).
		Where("entity_id = ?", req.EntityID).
		First(&notificationObject).Error; err != nil {
		log.Println("SOS, notification err: ", err)
	}

	if err := n.db.
		Transaction(func(tx *gorm.DB) error {
			if err := tx.
				Where("notification_object_id = ?", notificationObject.ID).
				Where("notifier_id = ?", req.NotifierID).
				Delete(&notificationEntities.Notification{}).Error; err != nil {
				return err
			}
			if err := tx.
				Where("notification_object_id = ?", notificationObject.ID).
				Where("actor_id = ?", req.ActorID).
				Delete(&notificationEntities.NotificationChange{}).Error; err != nil {
				return err
			}
			if err := tx.
				Delete(&notificationObject).Error; err != nil {
				return err
			}

			return nil
		}); err != nil {
		log.Println("DeleteNotification tx: ", err)
	}
}
