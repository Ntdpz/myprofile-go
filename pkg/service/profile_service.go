package service

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"myprofile-go/pkg/models"
)

func init() {
	envPath := filepath.Join("..", "..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println(".env file loaded successfully")
	}
}

func GetProfile() models.Profile {
	personalInformation := models.PersonalInformation{
		FullName:        os.Getenv("FULL_NAME"),
		Address:         os.Getenv("ADDRESS"),
		PhoneNumber:     os.Getenv("PHONE_NUMBER"),
		Email:           os.Getenv("EMAIL"),
		LinkedInProfile: os.Getenv("LINKEDIN_PROFILE"),
		GithubProfile:   os.Getenv("GITHUB_PROFILE"),
	}

	education := models.Education{
		Degree:      os.Getenv("EDUCATION_DEGREE"),
		Institution: os.Getenv("EDUCATION_INSTITUTION"),
		Duration:    os.Getenv("EDUCATION_DURATION"),
		Major:       os.Getenv("EDUCATION_MAJOR"),
		GPA:         os.Getenv("EDUCATION_GPA"),
	}

	workExperience := models.WorkExperience{
		CompanyName:      os.Getenv("WORK_COMPANY_NAME"),
		Position:         os.Getenv("WORK_POSITION"),
		Duration:         os.Getenv("WORK_DURATION"),
		Responsibilities: os.Getenv("WORK_RESPONSIBILITIES"),
		Achievements:     os.Getenv("WORK_ACHIEVEMENTS"),
	}

	skills := models.Skills{
		TechnicalSkills:     []string{os.Getenv("SKILLS_TECHNICAL")},
		CommunicationSkills: []string{os.Getenv("SKILLS_COMMUNICATION")},
		TeamworkSkills:      []string{os.Getenv("SKILLS_TEAMWORK")},
		CollaborationTools:  []string{os.Getenv("SKILLS_COLLABORATION")},
	}

	projects := []map[string]string{
		{
			"project_name":      os.Getenv("PROJECT_NAME"),
			"description":       os.Getenv("PROJECT_DESCRIPTION"),
			"technologies_used": os.Getenv("PROJECT_TECHNOLOGIES_USED"),
		},
	}

	languages := []models.Languages{
		{
			Language:    os.Getenv("LANGUAGE"),
			Proficiency: os.Getenv("LANGUAGE_PROFICIENCY"),
		},
	}

	// ส่งคืนข้อมูลโปรไฟล์
	return models.Profile{
		PersonalInformation: personalInformation,
		Education:           []models.Education{education},
		WorkExperience:      []models.WorkExperience{workExperience},
		Skills:              skills,
		Projects:            projects,
		Languages:           languages,
	}
}
