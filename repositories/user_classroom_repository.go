package repositories

import (
	"errors"
	"go_starter/models"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type UserClassRoomRepository interface {
	DeleteUserClassRoomAssociation(userID uint, classroomID uint) error
	CheckUserClassRoomExistsRepository(userID uint, classRoomID uint) (bool, error)
	CreateUserClassRoomRepository(request []models.UserClassroom) error
	GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error)

	//GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error)
	GetClassroomByTeacherRepository(userId int, userType string) ([]models.UserClassroom, error)

	GetAllUserClassRoomRepository() ([]models.UserClassroom, error)
	GetByIdUserClassRoomRepository(id uint) (*models.UserClassroom, error)
	GetByClassroomIdRepository(classRoomID uint) (*models.UserClassroom, error)
	GetByUserIDRepository(userID uint) (*models.UserClassroom, error)

	UpdateUserClassRoomRepository(request *models.UserClassroom) error
	DeleteUserClassRoomByUserIDRepository(userID uint) error
}

type userClassRoomRepository struct{ db *gorm.DB }

func (uc *userClassRoomRepository) DeleteUserClassRoomAssociation(userID uint, classroomID uint) error {
	result := uc.db.Where("user_id = ? AND classroom_id = ?", userID, classroomID).Delete(&models.UserClassroom{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *userClassRoomRepository) CheckUserClassRoomExistsRepository(userID uint, classRoomID uint) (bool, error) {
	var count int64
	if err := uc.db.Model(&models.UserClassroom{}).
		Where("user_id = ?", userID).
		Where("classroom_id = ?", classRoomID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (uc *userClassRoomRepository) CreateUserClassRoomRepository(request []models.UserClassroom) error {

	//if err := uc.db.Create(&request).Error; err != nil {
	//	return err
	//}
	//return nil

	if err := uc.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&request).Error; err != nil {
		return err
	}
	return nil
}

func (uc *userClassRoomRepository) GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error) {
	var userClassrooms []models.UserClassroom

	// Preload User and Classroom associations with conditions
	query := uc.db.
		Preload("User", "user_type = ?", userType).
		Preload("Classroom").
		Where("classroom_id = ?", classroomID).
		Find(&userClassrooms)

	if query.Error != nil {
		return nil, query.Error
	}

	return userClassrooms, nil
}

//func (uc *userClassRoomRepository) GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error) {
//	var model []models.UserClassroom
//
//	// Construct the raw SQL query
//	sqlQuery := `
//        SELECT user_classrooms.*, classrooms.*, users.* FROM user_classrooms
//        JOIN classrooms ON user_classrooms.classroom_id = classrooms.id
//        JOIN users ON user_classrooms.user_id = users.id
//        WHERE user_classrooms.classroom_id = ? AND users.user_type = ?
//    `
//
//	// Execute the raw SQL query
//	query := uc.db.Raw(sqlQuery, classroomID, userType).Scan(&model)
//	if query.Error != nil {
//		return nil, query.Error
//	}
//
//	return model, nil
//}

//func (uc *userClassRoomRepository) GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error) {
//	var model []models.UserClassroom
//	query := uc.db.Preload("Classroom").Preload("User").
//		Joins("JOIN users ON user_classrooms.user_id = users.id").
//		Where("user_classrooms.classroom_id = ? AND users.user_type = ?", classroomID, userType).
//		Find(&model).Error
//	if query != nil {
//		return nil, query
//	}
//	return model, nil
//}

//func (uc *userClassRoomRepository) GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error) {
//	var model []models.UserClassroom
//	query := uc.db.Preload("Classroom").Preload("User").Where("classroom_id=? AND user_type=?", classroomID, userType).Find(&model).Error
//	if query != nil {
//		return nil, query
//	}
//	return model, nil
//}

func (uc *userClassRoomRepository) GetClassroomByTeacherRepository(userId int, userType string) ([]models.UserClassroom, error) {
	var model []models.UserClassroom
	query := uc.db.Preload("ClassRoom").Preload("User").Where("user_id=? AND user_type=?", userId, userType).Find(&model).Error
	if query != nil {
		return nil, query
	}
	return model, nil
}

func (uc *userClassRoomRepository) DeleteUserClassRoomByUserIDRepository(userID uint) error {

	var count int64
	if err := uc.db.Model(&models.UserClassroom{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("user class room not found")
	}

	// Delete the UserClass
	if err := uc.db.Where("user = ?", userID).Delete(&models.UserClassroom{}).Error; err != nil {
		return err
	}
	return nil
}

func (uc *userClassRoomRepository) GetAllUserClassRoomRepository() ([]models.UserClassroom, error) {

	var model []models.UserClassroom
	query := uc.db.Find(&model).Error

	if query != nil {
		return nil, query
	}
	return model, nil
}

func (uc *userClassRoomRepository) GetByClassroomIdRepository(classRoomID uint) (*models.UserClassroom, error) {

	var model models.UserClassroom

	// Execute raw SQL query
	query := uc.db.Raw("SELECT * FROM user_class_rooms WHERE class_room_id = ?", classRoomID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (uc *userClassRoomRepository) GetByUserIDRepository(userID uint) (*models.UserClassroom, error) {

	var model models.UserClassroom

	// Execute raw SQL query
	query := uc.db.Raw("SELECT * FROM user_class_rooms WHERE user_id = ?", userID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (uc *userClassRoomRepository) GetByIdUserClassRoomRepository(id uint) (*models.UserClassroom, error) {

	var model models.UserClassroom

	// Execute raw SQL query
	query := uc.db.Raw("SELECT * FROM user_class_rooms WHERE id = ?", id).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (uc *userClassRoomRepository) UpdateUserClassRoomRepository(request *models.UserClassroom) error {

	query := uc.db.Model(&models.UserClassroom{}).Where("user_id = ?", request.UserID).Updates(request)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("no user_id found")
	}
	return nil
}

func NewUserClassRoom(db *gorm.DB) UserClassRoomRepository {
	//db.Migrator().DropTable(models.UserClassroom{}, models.User{}, models.ClassRoom{}, models.UserBehavior{})
	//db.AutoMigrate(models.UserClassroom{})

	return &userClassRoomRepository{db: db}
}
