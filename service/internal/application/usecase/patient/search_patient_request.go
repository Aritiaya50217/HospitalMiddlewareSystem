package patient

type SearchPatientRequest struct {
	PatientID   int64  `json:"patient_id"`
	NationalID  string `json:"national_id"`
	PassportID  string `json:"passport_id"`
	FirstName   string `form:"first_name"`
	MiddleName  string `form:"middle_name"`
	LastName    string `form:"last_name"`
	DateOfBirth string `form:"date_of_birth"` // YYYY-MM-DD
	PhoneNumber string `form:"phone_number"`
	Email       string `form:"email"`
}
