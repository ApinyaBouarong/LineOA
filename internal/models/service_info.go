package models

type ServiceInfo struct {
	PatientInfo      PatientInfo `json:"patient_info"`
	ID               int         `json:"id"`
	ServiceID        string      `json:"service_id"`
	ServiceCode      string      `json:"service_code"`
	RightToTreatment string      `json:"right_to_treatment"`
	ServiceType      string      `json:"service_type"`
	Activity         string      `json:"activity"`
	Location         string      `json:"location"`
	Day              string      `json:"day"`
	StartTime        string      `json:"start_time"`
	EndTime          string      `json:"end_time"`
	Period           string      `json:"period"`
	Selected         bool        `json:"selected"`
}
