package main

import "time"

type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Gender    string
	Password  string
	Role      string
	Active    bool
}

type Companies struct {
	CompanyID             int
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
	BranchID int
	Street   int
	City     string
	State    string
	Country  string
	Zip      int
}

type CompanyBranches struct {
	CompanyIDRef int
	BranchIDRef  int
}

type Regions struct {
	RegionID     int
	Region       string
	Geography    string
	RegionNotes  string
	CompanyIDRef int
}

type RegionBranches struct {
	RegionIDRef int
	BranchIDRef int
}

type DepartmentList struct {
	Department string
}

type JobInterviewQuestions struct {
	JobIDRef               int
	InterviewQuestionIDRef int
	IncludeFlag            bool
}

type Departments struct {
	DepartmentID int
	Department   string
	DeleteFlag   bool
	CompanyIDRef int
}

type Questions struct {
	QuestionID            int
	Question              string
	QuestionType          string
	QuestionCategoryIDRef int
	QuestionAnswerType    string
}

type JobDepartment struct {
	JobDepartmentID int
	JobIDRef        int
	DepartmentIDRef int
}

type Tasks struct {
	TaskID          int
	Task            string
	TaskDescription string
	TaskCategory    string
	DeleteFlag      bool
	CompanyIDRef    int
}

type Tasks1 struct {
	TaskID          int
	Task            string
	TaskDescription string
	TaskCategory    string
	DeleteFlag      bool
	CompanyIDRef    int
}

type CareerPaths struct {
	JobIDRef     int
	NextJobIDRef int
}

type Jobs struct {
	JobID          int
	Job            string
	JobSeries      string
	JobGrade       string
	JobDescription string
	SalaryLow      int
	SalaryHigh     int
	DeleteFlag     bool
	JobCode        int
}

type CourseLearningOutcomes struct {
	CourseIDRef          int
	LearningOutcomeIDRef int
	SortNumber           int
}

type EmployeeCompensation struct {
	EmployeeIDRef             int
	EmployeeSalary            int
	EmployeePayBasis          string
	EmployeePaySchedule       string
	EmployeePayCommissionFlag bool
	EmployeePayCommissionRate bool
	EmployeePayBonusFlag      bool
	EmployeePayBonusRate      int16
	EmployeePayBonusSchedule  string
}

type EmployeeAchievements struct {
	EmployeeIDRef int
	Achievement   string
	Date          time.Time
	Where         string
	Notes         string
}

type Subscribers struct {
	SubscriberName          string
	SubscriberPassword      string
	SubscriberCompanyIDRef  int
	SubscriberDate          time.Time
	SubscriberType          string
	SubscriberFee           int
	SubscriberPaid          bool
	SubscriberPaidDate      time.Time
	SubscriberExperienceRep string
}

type PeerCodes struct {
	PeerID      int
	PeerType    string
	RowSource   string
	ColumnCount int
	ListWidth   int
	Caption     string
}

type BusinessResults struct {
	BusinessResultsID   int
	BusinessReult       string
	BusinessResultNotes string
	DeleteFlag          bool
	CompanyIDRef        int
}

type EnvironmentChanges struct {
	EnvironmentChangeID int
	CompanyIDRef        int
	EmployeeIDRef       int
	EnvironmentChange   string
	EffectiveDate       time.Time
	ExpireDate          time.Time
	Notes               string
}

type DepartmentManagers struct {
	DepartmentIDRef  int
	MGREmployeeIDRef int
	EffectiveDate    time.Time
	DeleteFlag       bool
}

type EmployeePerfObservations struct {
	PerfObservationID   int
	EmployeeIDRef       int
	TaskIDRef           int
	ToBeAddressedFlag   bool
	PracticeRating      int
	PerfObservation     string
	PerfObservationDate time.Time
	Notes               string
	ScheduledReviewDate time.Time
}

type TaskSkills struct {
	TaskSkillID     int
	TaskIDRef       int
	SkillIDRef      int
	TaskSkillRating int
}

type TaskKnowledge struct {
	TaskKnowledgeID     int
	TaskIDRef           int
	KnowledgeIDRef      int
	TaskKnowledgeRating int
}

type TaskBehavior struct {
	TaskIDRef          int
	BehaviorIDRef      int
	TaskBehaviorRating int
}

type JobTasks struct {
	JobTaskID          int
	JobIDRef           int
	TaskIDRef          int
	JobTaskJobFunction string
	JobTaskPercent     int
	JobTaskNotes       string
}

type JobCources struct {
	JobIDRef    int
	CourseIDRef int
	Notes       string
}

type TaskProcedures struct {
	TaskProcedureID int
	TaskIDRef       int
	ProcedureIDRef  int
}

type ProceduresAndJobAids struct {
	ProcedureAndJobAidID   int
	BusinessReference      string
	ItemDescription        string
	ProcedureAndJobAidItem string
	CompanyIDRef           int
}

type JobTaskPerformanceMeasurements struct {
	JobTaskIDRef                       int
	PerformanceMeasurementIDRef        int
	PerformanceMeasurementJobTaskNotes string
}

type JobSuccessFactors struct {
	JobIDRef          int
	JobSuccessFactor  string
	SuccessEnablerHow string
}

type JobAccountabilityFactors struct {
	JobIDRef                           int
	JobAccountabilityFactor            string
	JobAccountabilityFactorMeasurement string
	JobAccountabilityFactorNotes       string
}

type CompenstationSchedule struct {
	CompensationScheduleID int
	CompansationSchedule   string
	DeleteFlag             bool
}

type TenureTypes struct {
	TenureID    int
	TenureType  string
	TenureStart time.Time
	TenureEnd   time.Time
}

type VacationDaysRules struct {
	WorkTenureDays       int
	EligibleVacationDays int
	Notes                string
	DeleteFlag           bool
}

type SickDaysRules struct {
	WorkTenureDays       int
	EligibleVacationDays int
	Notes                string
	DeleteFlag           bool
}

type PersonalDaysRules struct {
	WorkTenureDays       int
	EligiblePersonalDays int
	Notes                string
	DeleteFlag           bool
	CompanyIDRef         int
}

type PerformanceMeasurementTypes struct {
	PerformanceMeasurementTypeID       int
	PerformanceMeasurementType         string
	PerformanceMeasurementNotes        string
	DeleteFlag                         bool
	UpperThreshold                     int
	LowerThreshold                     int
	PerformanceMeasurementTypeCategory string
	CompanyIDRef                       int
}

type PerformanceMeasurements struct {
	CompanyIDRef               int
	PerformanceMeasurementType string
	EmployeeIDRef              int
	MeasurementDate            time.Time
	Measurement                string
	MeasurementNotes           string
}

type BuisinessResultsPerformanceMeasurementTypes struct {
	BuisinessResultsPerformanceMeasurementTypeID    int
	BusinessResultIDRef                             int
	PerformanceMeasurementIDRef                     int
	BuisinessResultsPerformanceMeasurementTypeHow   string
	BuisinessResultsPerformanceMeasurementTypeNotes string
}

type LocationsAndDirections struct {
	LocationID    int
	Location      string
	Directions    string
	LocationPhone string
	CompanyIDRef  int
}

type CourseKnowledge struct {
	CourseIDRef    int
	KnowledgeIDRef int
}

type CourseBehaviors struct {
	CourseIDRef   int
	BehaviorIDRef int
}

type EmployeeManagers struct {
	EmployeeIDRef int
	ManagerIDRef  int
	EffectiveDate time.Time
	DeleteFlag    bool
}

type EmployeeWatchLists struct {
	EmployeeIDRef  int
	WatchListIDRef int
}

type WatchLists struct {
	WatchListID   int
	WatchListName string
	WatchListInfo string
	CompanyIDRef  int
}
type EmployeeIssueGoals struct {
	IssueGoalID    int
	PerfObsrvIDRef int
	IssueGoal      string
	Notes          string
	EmployeeIDRef  int
	GoalSource     string
}

type TrainingVendors struct {
	VendorID int
	Vendor   string
}

type CourseTypes struct {
	CourseTypeID int
	CourseType   string
}

type Materials struct {
	MaterialID          int
	MaterialDescription string
	Material            string
	CompanyIDRef        int
}

type CourseFollowup struct {
	CourseIDRef   int
	FollowUpIDRef int
}

type StatusValues struct {
	StatusValue string
}

type CourseMaterials struct {
	CourseIDRef    int
	MaterialsIDRef int
	Quantity       int
}

type CoursePrework struct {
	CourseIDRef  int
	PreworkIDRef int
}

type Sessions struct {
	SessionID        int
	CourseIDRef      int
	InstructorIDRef  int
	Date             time.Time
	Time             int
	LocationIDRef    int
	Cost             int
	Notes            string
	ClassNumber      int
	ClassMeetingType string
	ClassLimit       int
	InstructorIDRef2 int
	InstructorIDRef3 int
	InstructorIDRef4 int
	InstructorIDRef5 int
	InstructorIDRef6 int
}
type Courses struct {
	CourseID               int
	CourseRef              string
	Category               string
	CourseTitle            string
	CourseDescription      string
	Notes                  string
	CourseType             string
	Vendor                 string
	VendorType             string
	ClassLimit             int
	DeleteFlag             bool
	IncludeInCourseCatalog bool
	CompanyIDRef           int
}

type Activities struct {
	ActivityID                    int
	CourseIDRef                   int
	SessionIDRef                  int
	EmployeeIDRef                 int
	LocationIDRef                 int
	Status                        string
	Feedback                      string
	ReferralEmplIDRef             int
	FacilitatorNotesOnParticipant string
}

type EmployeeAccountabilityReviews struct {
	AcctReviewID              int
	EmployeeIDRef             int
	AcctReviewDate            time.Time
	Notes                     string
	SchedReviewDate           time.Time
	ActReviewDate             time.Time
	ReviewedByEmplIDRef       int
	AcctFactor1               string
	AcctFactor1ScoreEmployee  string
	AcctFactor1ScoreManager   string
	AcctFactor1ScoreBoth      string
	AcctFactor1Action         string
	AcctFactor2               string
	AcctFactor2ScoreEmployee  string
	AcctFactor2ScoreManager   string
	AcctFactor2ScoreBoth      string
	AcctFactor2Action         string
	AcctFactor3               string
	AcctFactor3ScoreEmployee  string
	AcctFactor3ScoreManager   string
	AcctFactor3ScoreBoth      string
	AcctFactor3Action         string
	AcctFactor4               string
	AcctFactor4ScoreEmployee  string
	AcctFactor4ScoreManager   string
	AcctFactor4ScoreBoth      string
	AcctFactor4Action         string
	AcctFactor5               string
	AcctFactor5ScoreEmployee  string
	AcctFactor5ScoreManager   string
	AcctFactor5ScoreBoth      string
	AcctFactor5Action         string
	AcctFactor6               string
	AcctFactor6ScoreEmployee  string
	AcctFactor6ScoreManager   string
	AcctFactor6ScoreBoth      string
	AcctFactor6Action         string
	AcctFactor7               string
	AcctFactor7ScoreEmployee  string
	AcctFactor7ScoreManager   string
	AcctFactor7ScoreBoth      string
	AcctFactor7Action         string
	AcctFactor8               string
	AcctFactor8ScoreEmployee  string
	AcctFactor8ScoreManager   string
	AcctFactor8ScoreBoth      string
	AcctFactor8Action         string
	AcctFactor9               string
	AcctFactor9ScoreEmployee  string
	AcctFactor9ScoreManager   string
	AcctFactor9ScoreBoth      string
	AcctFactor9Action         string
	AcctFactor10              string
	AcctFactor10ScoreEmployee string
	AcctFactor10ScoreManager  string
	AcctFactor10ScoreBoth     string
	AcctFactor10Action        string
	AcctFactor11              string
	AcctFactor11ScoreEmployee string
	AcctFactor11ScoreManager  string
	AcctFactor11ScoreBoth     string
	AcctFactor11Action        string
	AcctFactor12              string
	AcctFactor12ScoreEmployee string
	AcctFactor12ScoreManager  string
	AcctFactor12ScoreBoth     string
	AcctFactor12Action        string
}

type CoursePrerequisites struct {
	CourseIDRef             int
	PrerequisiteCourseIDRef int
}

type JobDepartments1 struct {
	JobDepartmentID int
	JobIDRef        int
	DepartmentIDRef int
}

type Jobs1 struct {
	JobID          int
	Job            string
	JobSeries      string
	JobGrade       string
	JobDescription string
	SalaryLow      int
	SalaryHight    int
	DeleteFlag     bool
	JobCode        int
	CompanyIDRef   int
}

type Jobs2 struct {
	JobID          int
	Job            string
	JobSeries      string
	JobGrade       string
	JobDescription string
	SalaryLow      int
	SalaryHigh     int
	DeleteFlag     bool
	JobCode        int
	CompanyIDRef   int
}

type EmployeeInterviews struct {
	IntervieweeEmplIDRef int
	InterviewerEmplIDRef int
	LocationIDRef        int
	Date                 time.Time
	InterveiewPurpose    string
	DeleteFlag           bool
}

type InterviewPurposes struct {
	InterviewPurpose string
}

type PerformanceMeasurementTypes1 struct {
	PerformanceMeasurementTypeID       int
	PerformanceMeasurementType         string
	PerformanceMeasurementNotes        string
	DeleteFlag                         bool
	UpperThreshold                     string
	LowerThreshold                     string
	PerformanceMeasurementTypeCategory string
	CompanyIDRef                       int
	MeasurementType                    string
}

type EmployeePriorJobs struct {
	EmployeeIDRef int
	PriorCompany  string
	PriorJob      string
	StartDate     time.Time
	EndDate       time.Time
	Notes         string
}

type EmployeeContacts struct {
	EmployeeIDRef int
	ContactName   string
	ContactPhone  int
	ContactStreet string
	ContactCity   string
	ContactState  string
	ContactZip    int
	Primary       string
}

type EmployeeDaysOut struct {
	EmployeeIDRef int
	DayOutPurpose string
	DayOutDate    time.Time
	DeleteFlag    bool
}

type Schedules struct {
	ScheduleID int
	Schedule   string
	DeleteFlag bool
}

type PersonnelData struct {
	EmployeeID          int
	FullName            string
	FirstName           string
	LastName            string
	CompanyIDRef        int
	CompanyDate         int
	BranchIDRef         int
	BranchDate          time.Time
	HireDate            time.Time
	MGREmplIDRef        int
	MGRDate             time.Time
	DepartmentIDRef     int
	DepartmentDate      time.Time
	JobDepartmentIDRef  int
	JobDate             time.Time
	ScheduleIDRef       int
	ScheduleDate        time.Time
	Status              string
	StatusDate          time.Time
	MGRFlag             bool
	AdminFlag           bool
	InterviewerFlag     bool
	InstructorFlag      bool
	MiddleName          string
	Salutation          string
	Suffix              string
	EmployeeNumber      int
	Soc                 string
	WorkPhone1          int
	WorkPhone1Ext       int
	WorkPhone1Type      string
	WorkPhone2          int
	WorkPhone2Ext       int
	WorkPhone2Type      string
	WorkPhone3          int
	WorkPhone3Ext       int
	WorkPhone3Type      string
	WorkInternetAddress string
	HomeAddressLine1    string
	HomeAddressLine2    string
	HomeCity            string
	HomeState           string
	HomeZip             int
	HomePhone1          int
	HomePhone1Type      string
	HomeInternetAddress string
	DeleteFlag          bool
	Birthdate           time.Time
	Source              string
	DateAdded           time.Time
	EmployeeNotes       string
	EmployeeNotes2      string
}
type PersonnelData1 struct {
	EmployeeID          int
	FullName            string
	FirstName           string
	LastName            string
	CompanyIDRef        int
	CompanyDate         int
	BranchIDRef         int
	BranchDate          time.Time
	HireDate            time.Time
	MGREmplIDRef        int
	MGRDate             time.Time
	DepartmentIDRef     int
	DepartmentDate      time.Time
	JobDepartmentIDRef  int
	JobDate             time.Time
	ScheduleIDRef       int
	ScheduleDate        time.Time
	Status              string
	StatusDate          time.Time
	MGRFlag             bool
	AdminFlag           bool
	InterviewerFlag     bool
	InstructorFlag      bool
	MiddleName          string
	Salutation          string
	Suffix              string
	EmployeeNumber      int
	Soc                 string
	WorkPhone1          int
	WorkPhone1Ext       int
	WorkPhone1Type      string
	WorkPhone2          int
	WorkPhone2Ext       int
	WorkPhone2Type      string
	WorkPhone3          int
	WorkPhone3Ext       int
	WorkPhone3Type      string
	WorkInternetAddress string
	HomeAddressLine1    string
	HomeAddressLine2    string
	HomeCity            string
	HomeState           string
	HomeZip             int
	HomePhone1          int
	HomePhone1Type      string
	HomeInternetAddress string
	DeleteFlag          bool
	Birthdate           time.Time
	Source              string
	DateAdded           time.Time
	EmployeeNotes       string
	EmployeeNotes2      string
}

type PersonnelData2 struct {
	EmployeeID          int
	FullName            string
	FirstName           string
	LastName            string
	CompanyIDRef        int
	CompanyDate         int
	BranchIDRef         int
	BranchDate          time.Time
	HireDate            time.Time
	MGREmplIDRef        int
	MGRDate             time.Time
	DepartmentIDRef     int
	DepartmentDate      time.Time
	JobDepartmentIDRef  int
	JobDate             time.Time
	ScheduleIDRef       int
	ScheduleDate        time.Time
	Status              string
	StatusDate          time.Time
	MGRFlag             bool
	AdminFlag           bool
	InterviewerFlag     bool
	InstructorFlag      bool
	MiddleName          string
	Salutation          string
	Suffix              string
	EmployeeNumber      int
	Soc                 string
	WorkPhone1          int
	WorkPhone1Ext       int
	WorkPhone1Type      string
	WorkPhone2          int
	WorkPhone2Ext       int
	WorkPhone2Type      string
	WorkPhone3          int
	WorkPhone3Ext       int
	WorkPhone3Type      string
	WorkInternetAddress string
	HomeAddressLine1    string
	HomeAddressLine2    string
	HomeCity            string
	HomeState           string
	HomeZip             int
	HomePhone1          int
	HomePhone1Type      string
	HomeInternetAddress string
	DeleteFlag          bool
	Birthdate           time.Time
	Source              string
	DateAdded           time.Time
	EmployeeNotes       string
	EmployeeNotes2      string
}

type EmployeeBehaviors struct {
	EmployeeIDRef int
	BehaviorIDRef int
}

type EmployeeSchedules struct {
	EmployeeIDRef int
	ScheduleIDRef int
	EffectiveDate time.Time
	DeleteFlag    bool
}

type EmployeeJobs struct {
	EmployeeIDRef      int
	JobDepartmentIDRef int
	EffectiveDate      time.Time
	DeleteFlag         bool
}

type EmployeeDepartments struct {
	EmployeeIDRef   int
	DepartmentIDRef int
	EffectiveDate   time.Time
	DeleteFlag      bool
}

type EmployeeSkilss struct {
	EmployeeIDRef int
	SkillIDRef    int
}
type EmployeeKnowledge struct {
	EmployeeIDRef  int
	KnowledgeIDRef int
}

type EmployeeCompanies struct {
	EmployeeIDRef int
	CompanyIDRef  int
	EffectiveDate time.Time
	DeleteFlag    bool
}

type EmployeeBranches struct {
	EmployeeIDRef int
	BranchIDRef   int
	EffectiveDate time.Time
	DeleteFlag    bool
}

type EmployeeGoalActions struct {
	GoalActionID    int
	IssueGoalIDRef  int
	EstStartDate    time.Time
	ActStartDate    time.Time
	EstCompleteDate time.Time
	ActCompleteDate time.Time
	GoalAction      string
	Notes           string
	EmployeeIDRef   int
}
