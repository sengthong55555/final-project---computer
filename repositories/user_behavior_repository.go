package repositories

import (
	"errors"
	"go_starter/models"
	"gorm.io/gorm"
	"time"
)

type UserBehaviorRepository interface {
	UpdateUserBehaviorRepository(request []models.UserBehavior) error
	InsertUserBehaviorRepository(request []models.UserBehavior) error

	GetAllUserBehaviorRepository() ([]models.UserBehavior, error)
	GetByIdUserBehaviorRepository(id uint) (*models.UserBehavior, error)
	GetByClassroomIdRepository(classRoomID uint) (*models.UserBehavior, error)
	GetByUserIDRepository(userID uint) (*models.UserBehavior, error)
	CreateUserBehaviorRepository(request *models.UserBehavior) error

	DeleteUserBehaviorByUserIDRepository(userID uint) error
}

type userBehaviorRepository struct{ db *gorm.DB }

func (ub *userBehaviorRepository) UpdateUserBehaviorRepository(requests []models.UserBehavior) error {
	return ub.db.Transaction(func(tx *gorm.DB) error {
		for _, request := range requests {
			var model models.UserBehavior

			// Find the specific user behavior record by userID and classroomID
			if err := tx.Where("user_id = ? AND classroom_id = ?", request.UserID, request.ClassroomID).
				First(&model).Error; err != nil {
				return err
			}

			// Increment count fields based on request
			if request.StudentCheck {
				model.CountCheck++
			}
			if request.StudentAbsent {
				model.CountAbsent++
			}
			if request.StudentVacation {
				model.CountVacation++
			}
			if request.StudentBreakRule {
				model.CountBreakRule++
			}
			model.UpdatedAt = time.Now()

			// Save the updated user behavior record into the database
			if err := tx.Save(&model).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

//func (ub *userBehaviorRepository) UpdateUserBehaviorRepository(requests []models.UserBehavior) error {
//	return ub.db.Transaction(func(tx *gorm.DB) error {
//		for _, request := range requests {
//			var model models.UserBehavior
//
//			// Find the specific user behavior record by userID and classroomID
//			if err := tx.Where("user_id = ? AND classroom_id = ?", request.UserID, request.ClassroomID).
//				First(&model).Error; err != nil {
//				return err
//			}
//
//			// Check and update statuses and increment count fields
//			if request.StudentCheck && !model.StudentCheck {
//				model.StudentCheck = true
//				model.CountCheck++
//			}
//			if request.StudentAbsent && !model.StudentAbsent {
//				model.StudentAbsent = true
//				model.CountAbsent++
//			}
//			if request.StudentVacation && !model.StudentVacation {
//				model.StudentVacation = true
//				model.CountVacation++
//			}
//			if request.StudentBreakRule && !model.StudentBreakRule {
//				model.StudentBreakRule = true
//				model.CountBreakRule++
//			}
//			model.UpdatedAt = time.Now()
//
//			// Save the updated user behavior record into the database
//			if err := tx.Save(&model).Error; err != nil {
//				return err
//			}
//		}
//		return nil
//	})
//}

func (ub *userBehaviorRepository) InsertUserBehaviorRepository(requests []models.UserBehavior) error {
	// Ensure ID is not manually set to avoid duplicate primary key issues
	for i := range requests {
		requests[i].ID = 0
		// Default values for the other fields are already set in the service
	}

	// Use bulk insert for better performance
	if err := ub.db.Create(&requests).Error; err != nil {
		return err
	}
	return nil
}

//func (u *userBehaviorRepository) InsertUserBehaviorRepository(request []models.UserBehavior) error {
//
//	// Ensure ID is not manually set to avoid duplicate primary key issues
//	for i := range request {
//		request[i].ID = 0
//		// Set all default student statuses to false
//		request[i].StudentCheck = false
//		request[i].StudentCheck = false
//		request[i].StudentAbsent = false
//		request[i].StudentVacation = false
//		request[i].StudentBreakRule = false
//		request[i].CountCheck = 0
//		request[i].CountAbsent = 0
//		request[i].CountVacation = 0
//		request[i].CountBreakRule = 0
//		request[i].CreatedAt = time.Now()
//		request[i].UpdatedAt = time.Now()
//	}
//
//	// Use bulk insert for better performance
//	if err := u.db.Create(&request).Error; err != nil {
//		return err
//	}
//	return nil
//}

func (ub *userBehaviorRepository) CreateUserBehaviorRepository(request *models.UserBehavior) error {

	if err := ub.db.Create(request).Error; err != nil {
		return err
	}
	return nil
}

func (ub *userBehaviorRepository) DeleteUserBehaviorByUserIDRepository(userID uint) error {

	var count int64
	if err := ub.db.Model(&models.UserClassroom{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("user class room not found")
	}

	// Delete the UserClass
	if err := ub.db.Where("user = ?", userID).Delete(&models.UserClassroom{}).Error; err != nil {
		return err
	}
	return nil
}

func (ub *userBehaviorRepository) GetAllUserBehaviorRepository() ([]models.UserBehavior, error) {

	var model []models.UserBehavior
	query := ub.db.Find(&model).Error

	if query != nil {
		return nil, query
	}
	return model, nil
}

func (ub *userBehaviorRepository) GetByClassroomIdRepository(classRoomID uint) (*models.UserBehavior, error) {

	var model models.UserBehavior

	// Execute raw SQL query
	query := ub.db.Raw("SELECT * FROM user_behaviors WHERE classroom_id = ?", classRoomID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (ub *userBehaviorRepository) GetByIdUserBehaviorRepository(id uint) (*models.UserBehavior, error) {

	var model models.UserBehavior

	// Execute raw SQL query
	query := ub.db.Raw("SELECT * FROM user_behaviors WHERE id = ?", id).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (ub *userBehaviorRepository) GetByUserIDRepository(userID uint) (*models.UserBehavior, error) {

	var model models.UserBehavior

	// Execute raw SQL query
	query := ub.db.Raw("SELECT * FROM user_behaviors WHERE user_id = ?", userID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func NewUserBehaviorRepository(db *gorm.DB) UserBehaviorRepository {
	//db.Migrator().DropTable(models.UserBehavior{})
	//db.AutoMigrate(models.UserBehavior{})
	return &userBehaviorRepository{db: db}
}
