package models

import "time"

type UserRegisted struct {
	Initials   string `db:"initials"`
	Lastname    string `db:"lastname"`
	ID_number   string `db:"id_number"`
	Association string `db:"association"`
	QRCode 		string `db:"qrcode"`
}

type UserRegistedResponse struct{

}

type GetUserDetails struct {
// agm.get_users
	ID_number   string `db:"_id_number"`
}

type GetUserByDetails struct{
	ID   int `db:"id"`
	Organization  string `db:"organization"`	
	QRCODE  string `db:"qrcode_"`

}
type AssignUserToQECode struct{
	ID_number   string `db:"id"`
	Association string `db:"associations"`
	QRCode string `db:"QRCode_"`
}

type CheckInUser struct {
	ID          string    `db:"id_"`
	Link          string    `db:"link"`
}
type CheckInUserReturn struct{
	ID          string    `db:"id_number_"`
}
type UserDisplayed struct{
	Message string `db:""message"`
}

/*   _id uuid,
_firstname text,
surname text,
id_number integer
*/
type AllUser struct {
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	ID_number string `db:"id_number"`
}
type GetAllType struct {
}

//public.get_user_json
type GetUser struct {
	Firstname     string `db:"initials"`
	Lastname      string `db:"lastname"`
	CheckedInDate time.Time `db:"checked_in_date"`
	Picturelink   string `db:"picture_link"`
	IDNumber      string `db:"id_number"`
	Organization  string `db:"organization"`
}
type Login struct {
	ID_number string `json:"id_number"`
}

type EmptyGetter struct {
	FucntionName string `json:"fucntion_name"`
}

type Associations struct{
	AssociationName string `db:"a_name"`
	Arrived int `db:"arrived"`
	Expected int `db:"expected"`
	ImageURL string `db:"image_url"`
}



type GetAssociations struct{
	Associations_Name string `db:"a_name_"`
	Expected int  `db:"expected_"`
	Arrived int `db:"arrived_"`
}

type MarkAsDisplayed struct{
	ID_Number string `db:"_id_number"`
}