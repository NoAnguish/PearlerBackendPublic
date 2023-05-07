package account

type Account struct {
	Id          string
	Name        string
	Email       string
	ImageURL    string `db:"image_url"`
	FirebaseUId string `db:"firebase_uid"`
	Description string `db:"description"`
}

type CreateAccountRequest struct {
	Email       string `json:"email,omitempty"`
	FirebaseUId string `json:"firebase_uid,omitempty"`
	Name        string `json:"name,omitempty"`
}

type AccountIdResponse struct {
	Id string `json:"id"`
}

type GetByFirebaseUIdResponseMetadata struct {
	Exists bool `json:"exists"`
}

type GetByFirebaseUIdResponse struct {
	Id       string                           `json:"id,omitempty"`
	Metadata GetByFirebaseUIdResponseMetadata `json:"metadata"`
}

type FirebaseUIdRequest struct {
	FirebaseUId string `json:"firebase_uid,omitempty"`
}

type AccountIdRequest struct {
	Id string `json:"id,omitempty"`
}

type AccountResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	ImageURL    string `json:"image_url"`
	FirebaseUId string `json:"firebase_uid"`
	Description string `json:"description"`
}

type AccountsListResponse struct {
	Accounts []AccountResponse `json:"accounts"`
}

type AccountExistsResponse struct {
	Exists bool `json:"exists"`
}

type UpdateAccountRequest struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
