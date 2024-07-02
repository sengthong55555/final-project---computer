package requests

type UserClassroomRequest struct {
	ClassroomID uint `json:"classroom_id" validate:"required"`
	//UserType    string `json:"user_type" validate:"required"`
}

type TeacherIdRequest struct {
	UserID   int    `json:"user_id"`
	UserType string `json:"user_type"`
}
type UserClassRoomRequest struct {
	ClassroomID uint   `json:"classroom_id"`
	UserIDs     []uint `json:"user_ids"`
}

type CreateUserClassRoomRequest struct {
	UserID      uint `json:"user_id"`
	ClassRoomID uint `json:"class_room_id"`
}

type UpdateUserClassRoomRequest struct {
	UserID      uint `json:"User_id"`
	ClassRoomID uint `json:"class_room_id"`
}
type DeleteUserClassRoomRequest struct {
	Id int `json:"id" validate:"required"`
}

type UserClassRoomIDRequest struct {
	Id int `json:"id"`
}
type UserClassRoomByIDRequest struct {
	ClassRoomID uint `json:"class_room_id"`
}
type UserClassRoomByUserIDRequest struct {
	UserID uint `json:"User_id"`
}
