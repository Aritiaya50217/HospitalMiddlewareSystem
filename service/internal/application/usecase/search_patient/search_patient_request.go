package search_patient

type SearchPatientRequest struct {
	PatientID  *int64  `json:"patient_id"`
	NationalID *string `json:"national_id"`
	PassportID *string `json:"passport_id"`
}
