package input

type UpdateDeviceInput struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}
