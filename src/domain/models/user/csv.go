package user

type ImportField struct {
	EmployeeID          string `csv:"employee_id"`
	FullName            string `csv:"fullname"`
	Barcode             string `csv:"barcode"`
	Organization        string `csv:"organization"`
	LocationId          int64  `csv:"location_id"`
	JobPosition         string `csv:"job_position"`
	JobLevel            string `csv:"job_level"`
	PositionLevelId     int64  `csv:"position_level_id"`
	JoinDate            string `csv:"join_date"`
	ResignDate          string `csv:"resign_date"`
	StatusEmployee      string `csv:"status"`
	EndDate             string `csv:"end_date"`
	SignDate            string `csv:"sign_date"`
	Email               string `csv:"email"`
	BirthDate           string `csv:"birthdate"`
	BirthPlace          string `csv:"birthplace"`
	CitizenIDAddress    string `csv:"citizen_id_address"`
	ResidentialAddress  string `csv:"residential_address"`
	NPWP                string `csv:"npwp"`
	PTKPStatus          string `csv:"ptkp_status"`
	EmployeeTaxStatus   string `csv:"employee_tax_status"`
	TaxConfig           string `csv:"tax_config"`
	BankName            string `csv:"bank_name"`
	BankAccount         string `csv:"bank_account"`
	BankAccountHolder   string `csv:"bank_account_holder"`
	BPJSKetenagakerjaan string `csv:"bpjs_ketenagakerjaan"`
	BPJSKesehatan       string `csv:"bpjs_kesehatan"`
	CitizenId           string `csv:"citizen_id"`
	MobilePhone         string `csv:"no_hp"`
	Phone               string `csv:"telp_rumah"`
	BranchName          string `csv:"branch_name"`
	Religion            string `csv:"religion"`
	Gender              string `csv:"gender"`
	MaritalStatus       string `csv:"marital_status"`
	NationalityCode     string `csv:"nationality_code"`
	Currency            string `csv:"currency"`
	LengthOfService     string `csv:"length_of_service"`
	PaymentSchedule     string `csv:"payment_schedule"`
	ApprovalLine        string `csv:"approval_line"`
	Manager             string `csv:"manager"`
	Grade               string `csv:"grade"`
	Class               string `csv:"class"`
	ProfilePicture      string `csv:"profile_picture"`
	GroupFFI            string `csv:"group_ffi"`
	CompanyIdentity     string `csv:"company_identity"`
	KTANumber           string `csv:"no_kta"`
	GadaNo              string `csv:"no_gada"`
	KTAExpiredDate      string `csv:"kta_expired_date"`
	AsalSekolah         string `csv:"asal_sekolah"`
	TahunLulus          int    `csv:"tahun_lulus"`
}
