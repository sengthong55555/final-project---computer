package requests

type CreateClassroomSubjectRequest struct {
	ClassroomID uint `json:"classroom_id"`
	SubjectID   uint `json:"subject_id"`
}
