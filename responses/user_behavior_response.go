package responses

type UserBehaviorResponse struct {
	ID                  uint   `json:"id" gorm:"primaryKey"`
	UserID              uint   `json:"user_id"`
	ClassRoomID         uint   `json:"class_room_id"`
	StudentCheck        bool   `json:"student_check"`
	StudentAbsent       bool   `json:"student_absent"`
	StudentVacation     bool   `json:"student_vacation"`
	StudentBreakingRule bool   `json:"student_breaking_rule"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

type MessageUserBehaviorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
type ReportChartMembershipResponse struct {
	Xaxis  Xaxis    `json:"xAxis"`
	Yaxis  Yaxis    `json:"yAxis"`
	Series []Series `json:"series"`
}

type Xaxis struct {
	Date []string `json:"date"`
}

type Yaxis struct {
	Number []string `json:"number"`
}

type Series struct {
	Name  string    `json:"name"`
	Color string    `json:"color"`
	Data  []float64 `json:"data"`
}
