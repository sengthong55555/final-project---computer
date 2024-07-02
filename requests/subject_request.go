package requests

type SubjectRequest struct {
	SubjectCode string `json:"subject_code" validate:"required"`
}

type InsertSubjectRequest struct {
	SubjectCode string `json:"subject_code" validate:"required"`
	SubjectName string `json:"subject_name" validate:"required"`
}
