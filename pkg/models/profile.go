package models

// PersonalInformation struct สำหรับข้อมูลส่วนตัว
type PersonalInformation struct {
	FullName        string `json:"full_name"`
	Address         string `json:"address"`
	PhoneNumber     string `json:"phone_number"`
	Email           string `json:"email"`
	LinkedInProfile string `json:"linkedin_profile"`
	GithubProfile   string `json:"github_profile"`
}

// Education struct สำหรับข้อมูลการศึกษา
type Education struct {
	Degree      string `json:"degree"`
	Institution string `json:"institution"`
	Duration    string `json:"duration"`
	Major       string `json:"major"`
	GPA         string `json:"gpa"`
}

// WorkExperience struct สำหรับข้อมูลประสบการณ์การทำงาน
type WorkExperience struct {
	CompanyName      string `json:"company_name"`
	Position         string `json:"position"`
	Duration         string `json:"duration"`
	Responsibilities string `json:"responsibilities"`
	Achievements     string `json:"achievements"`
}

// Skills struct สำหรับข้อมูลทักษะ
type Skills struct {
	TechnicalSkills     []string `json:"technical_skills"`
	CommunicationSkills []string `json:"communication_skills"`
	TeamworkSkills      []string `json:"teamwork_skills"`
	CollaborationTools  []string `json:"collaboration_tools"`
}

// Languages struct สำหรับข้อมูลภาษา
type Languages struct {
	Language    string `json:"language"`
	Proficiency string `json:"proficiency"`
}

// Profile struct รวมข้อมูลทั้งหมด
type Profile struct {
	PersonalInformation PersonalInformation `json:"personal_information"`
	Education           []Education         `json:"education"`
	WorkExperience      []WorkExperience    `json:"work_experience"`
	Skills              Skills              `json:"skills"`
	Projects            []map[string]string `json:"projects"`
	Languages           []Languages         `json:"languages"`
	ImageBase64         string              `json:"image_base64"`
}
