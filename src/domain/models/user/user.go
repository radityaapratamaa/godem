package user

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/kodekoding/phastos/v2/go/database"
	"github.com/volatiletech/null"
)

type (
	DataCommon struct {
		LoginRequest
		Name             string        `json:"name" db:"name"`
		PhoneNumber      *PhoneNumber  `json:"phone_number" db:"phone_number" col:"json"`
		BirthPlace       null.String   `json:"birthplace" db:"birthplace"`
		BirthDate        null.String   `json:"birthdate" db:"birthdate"`
		Document         *Documents    `json:"document" db:"document" col:"json"`
		MaritalStatus    null.String   `json:"marital_status" db:"marital_status"`
		BloodType        null.String   `json:"blood_type" db:"blood_type"`
		Religion         null.String   `json:"religion" db:"religion"`
		Gender           null.String   `json:"gender" db:"gender"`
		FamilyData       FamilyData    `json:"family_data" db:"family_data" col:"json"`
		EmergencyContact null.String   `json:"emergency_contact" db:"emergency_contact"`
		EducationData    EducationData `json:"education_data" db:"education_data" col:"json"`
		ProfilePic       null.String   `json:"profile_pic" db:"profile_pic"`
	}
	Data struct {
		database.BaseColumn[int64]
		DataCommon
		Address             *Address    `json:"address" db:"address" col:"json"`
		ActivationCode      null.String `json:"activation_code" db:"activation_code"`
		ActivationExpiredAt null.String `json:"activation_expired_at" db:"activation_expired_at"`
		ActiveAt            null.String `json:"active_at" db:"active_at"`
		UserGroupId         null.Int    `json:"user_group_id" db:"user_group_id"`
	}

	ResetDeviceRequest struct {
		UserId int64 `json:"user_id" schema:"user_id" validate:"required"`
	}

	Address struct {
		CitizenIDAddress string      `json:"alamat_ktp,omitempty" db:"alamat_ktp"`
		ResidentAddress  null.String `json:"alamat_domisili,omitempty" db:"alamat_domisili"`
	}

	// FamilyData will store map of relationship family data
	// ex: {
	//   "mother": {
	// 	     "name": "xxx",
	//       "phone_number": "0812xxxxx"
	// 	 },
	//   "father": {
	//       "name": "xxx",
	//	     "phone_number": "0812xxxxx"
	//   },
	//   .........
	//}
	FamilyData map[string]FamilyDetailData

	FamilyDetailData struct {
		Name        null.String `json:"name" db:"name"`
		PhoneNumber null.String `json:"phone_number" db:"phone_number"`
	}

	// EducationData will store map of level education data
	// ex: {
	//   "sd": {
	// 	     "name": "SD xxx",
	//       "graduate_year": 2000,
	//       "final_score": 8.6
	// 	 },
	//   "smp": {
	//       "name": "SMP xxx",
	//	     "graduate_year": 2003,
	//       "final_score": 9
	//   },
	//   .........,
	//   "universitas": {
	//       "name": "univ xxx",
	//	     "graduate_year": 2012,
	//       "final_score": 3.86
	//   },
	//}
	EducationData map[string]EducationDetailData

	EducationDetailData struct {
		Name         null.String  `json:"name" db:"name"`
		GraduateYear null.Int     `json:"graduate_year" db:"graduate_year"`
		FinalScore   null.Float32 `json:"final_score" db:"final_score"`
	}
)

// Value Make the address struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a Address) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Make the address struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *Address) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Value Make the address struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a FamilyData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Make the address struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *FamilyData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Value Make the address struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a EducationData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Make the address struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *EducationData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type PhoneNumber struct {
	MobilePhone string      `json:"no_hp,omitempty"`
	Phone       null.String `json:"telp_rumah,omitempty"`
}

// Value Make the phoneNumber struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a PhoneNumber) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Make the phoneNumber struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *PhoneNumber) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Documents struct {
	KTP             null.String `json:"ktp,omitempty" db:"ktp"`
	NPWP            null.String `json:"npwp,omitempty" db:"npwp"`
	CompanyIdentity null.String `json:"company_identity,omitempty" db:"company_identity"`
	KTA             KTADocument `json:"kta,omitempty" db:"kta"`
}

type KTADocument struct {
	No          null.String `json:"no,omitempty"`
	ExpiredDate null.String `json:"expired_date,omitempty"`
}

// Value Make the phoneNumber struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a Documents) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Make the phoneNumber struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *Documents) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Value Make the phoneNumber struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a KTADocument) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Make the phoneNumber struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *KTADocument) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Request struct {
	database.TableRequest
}
