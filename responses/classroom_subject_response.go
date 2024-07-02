package responses

type CreateClassRoomSubjectResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type ClassroomSubjectResponse struct {
	ID          uint   `json:"id"`
	ClassName   string `json:"class_name"`
	SubjectCode string `json:"subject_code"`
	SubjectName string `json:"subject_name"`
}
