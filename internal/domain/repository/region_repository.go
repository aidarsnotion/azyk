package repository

import (
	"azyk/internal/domain/models"
	"azyk/util/logger"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type RegionRepository interface {
	Create(region *models.Region) error
	GetByID(id int32) (*models.Region, error)
	List() ([]models.Region, error)
	Update(region *models.Region) (*models.Region, error)
	Delete(id int32) (*models.Region, error)
}

type regionRepository struct {
	db *gorm.DB
}

func NewRegionRepository(db *gorm.DB) RegionRepository {
	return &regionRepository{db: db}
}

func (r *regionRepository) Create(region *models.Region) error {
	if err := r.db.Create(region).Error; err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"error":     err.Error(),
			"operation": "create region",
		}).Error("Failed to create region")
		return fmt.Errorf("failed to create region: %w", err)
	}
	return nil
}

func (r *regionRepository) GetByID(id int32) (*models.Region, error) {
	var region models.Region
	if err := r.db.Preload("Parent").
		Preload("Children").
		Preload("Categories").
		Preload("Translations").
		First(&region, id).Error; err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"error":     err.Error(),
			"operation": "get region by id",
			"region_id": id,
		}).Error("Failed to get region by ID")
		return nil, fmt.Errorf("failed to get region by ID: %w", err)
	}
	return &region, nil
}

func (r *regionRepository) List() ([]models.Region, error) {
	var regions []models.Region
	if err := r.db.Preload("Parent").
		Preload("Children").
		Preload("Categories").
		Preload("Translations").
		Find(&regions).Error; err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"error":     err.Error(),
			"operation": "list regions",
		}).Error("Failed to list regions")
		return nil, fmt.Errorf("failed to list regions: %w", err)
	}
	return regions, nil
}

func (r *regionRepository) Update(region *models.Region) (*models.Region, error) {
	var existing models.Region
	if err := r.db.Preload("Translations").First(&existing, region.ID).Error; err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"error":     err.Error(),
			"operation": "find existing region before update",
			"region_id": region.ID,
		}).Error("Failed to find region for update")
		return nil, fmt.Errorf("failed to find existing region: %w", err)
	}

	region.CreatedAt = existing.CreatedAt

	if err := r.db.Model(&existing).
		Omit("CreatedAt").
		Updates(region).Error; err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"error":     err.Error(),
			"operation": "update region",
			"region_id": region.ID,
		}).Error("Failed to update region")
		return nil, fmt.Errorf("failed to update region: %w", err)
	}

	for _, newTranslation := range region.Translations {
		var existingTranslation models.RegionTranslation
		err := r.db.Where("region_id = ? AND lang = ?", region.ID, newTranslation.Lang).
			First(&existingTranslation).Error

		if err == nil {
			if err := r.db.Model(&existingTranslation).Updates(map[string]interface{}{
				"name":       newTranslation.Name,
				"updated_at": time.Now(),
			}).Error; err != nil {
				logger.Log.WithFields(map[string]interface{}{
					"error":     err.Error(),
					"operation": "update region translation",
					"region_id": region.ID,
					"lang":      newTranslation.Lang,
				}).Error("Failed to update region translation")
				return nil, fmt.Errorf("failed to update translation: %w", err)
			}
		} else if err == gorm.ErrRecordNotFound {
			newTranslation.RegionID = region.ID
			if err := r.db.Create(&newTranslation).Error; err != nil {
				logger.Log.WithFields(map[string]interface{}{
					"error":     err.Error(),
					"operation": "create region translation",
					"region_id": region.ID,
					"lang":      newTranslation.Lang,
				}).Error("Failed to create region translation")
				return nil, fmt.Errorf("failed to create new translation: %w", err)
			}
		} else {
			logger.Log.WithFields(map[string]interface{}{
				"error":     err.Error(),
				"operation": "check translation existence",
				"region_id": region.ID,
				"lang":      newTranslation.Lang,
			}).Error("Failed to check region translation existence")
			return nil, fmt.Errorf("error checking translation existence: %w", err)
		}
	}

	var updated models.Region
	if err := r.db.Preload("Translations").First(&updated, region.ID).Error; err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"error":     err.Error(),
			"operation": "load updated region",
			"region_id": region.ID,
		}).Error("Failed to load updated region")
		return nil, fmt.Errorf("failed to load updated region: %w", err)
	}

	return &updated, nil
}

func (r *regionRepository) Delete(id int32) (*models.Region, error) {
	tx := r.db.Begin()

	defer func() {
		if rec := recover(); rec != nil {
			logger.Log.WithFields(map[string]interface{}{
				"region_id": id,
				"panic":     rec,
			}).Error("Panic recovered during region deletion")
			tx.Rollback()
		}
	}()

	var region models.Region
	if err := tx.Preload("Translations").First(&region, id).Error; err != nil {
		tx.Rollback()
		logger.Log.WithFields(map[string]interface{}{
			"region_id": id,
			"error":     err.Error(),
		}).Error("Failed to find region before delete")
		return nil, fmt.Errorf("failed to find region before delete: %w", err)
	}

	//Мягко удаляем переводы/ хрен его пока знает не получается проставить deleted_at и передать клиенту в ответе
	if err := tx.Where("region_id = ?", id).Delete(&models.RegionTranslation{}).Error; err != nil {
		tx.Rollback()
		logger.Log.WithFields(map[string]interface{}{
			"region_id": id,
			"error":     err.Error(),
		}).Error("Failed to soft delete region translations")
		return nil, fmt.Errorf("failed to soft delete region translations: %w", err)
	}

	//Мягко удаляем сам регион
	if err := tx.Delete(&region).Error; err != nil {
		tx.Rollback()
		logger.Log.WithFields(map[string]interface{}{
			"region_id": id,
			"error":     err.Error(),
		}).Error("Failed to soft delete region")
		return nil, fmt.Errorf("failed to soft delete region: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"region_id": id,
			"error":     err.Error(),
		}).Error("Failed to commit transaction after soft delete")
		return nil, fmt.Errorf("failed to commit delete transaction: %w", err)
	}

	logger.Log.WithFields(map[string]interface{}{
		"region_id":   id,
		"region_code": region.Code,
		"region_name": region.Name,
	}).Info("Region and translations soft deleted successfully")

	return &region, nil
}
