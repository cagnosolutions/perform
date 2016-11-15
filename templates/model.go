package main

import "time"

type Companies struct {
	CompanyID             int64
	CompanyName           string
	DUNSNumber            int
	RegisteredDate        time.Time
	RegistrationFee       int
	RegistrationPaid      bool
	RegistrationPaidDate  time.Time
	CustomerExperienceRep string
	SalesRep              string
}

type Branches struct {
	BranchID int64
	Street   int
	City     string
	State    string
	Country  string
	Zip      int
}

type CompanyBranches struct {
	CompanyIDRef int64
	BranchIDRef  int64
}

type Regions struct {
	RegionID     int64
	Region       string
	Geography    string
	RegionNotes  string
	CompanyIDRef int64
}

type RegionBranches struct {
	RegionIDRef int64
	BranchIDRef int64
}

type DepartmentList struct {
	Department string
}

type JobInterviewQuestions struct {
	JobIDRef               int64
	InterviewQuestionIDRef int64
	IncludeFlag            bool
}

type Departments struct {
	DepartmentID int64
	Department   string
	DeleteFlag   bool
	CompanyIDRef int64
}

type Questions struct {
	QuestionID            int64
	Question              string
	QuestionType          string
	QuestionCategoryIDRef int64
	QuestionAnswerType    string
}

type JobDepartment struct {
	JobDepartmentID int64
	JobIDRef        int64
	DepartmentIDRef int64
}

type Tasks struct {
	TaskID          int64
	Task            string
	TaskDescription string
	TaskCategory    string
	DeleteFlag      bool
	CompanyIDRef    int64
}

type CareerPaths struct {
	JobIDRef     int64
	NextJobIDRef int64
}

type Jobs struct {
	JobID          int64
	Job            string
	JobSeries      string
	JobGrade       string
	JobDescription string
	SalaryLow      int
	SalaryHigh     int
	DeleteFlag     bool
	JobCode        int64
}

type CourseLearningOutcomes struct {
	CourseIDRef          int64
	LearningOutcomeIDRef int64
	SortNumber           int
}
