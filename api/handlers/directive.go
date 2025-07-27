package handlers

import (
	"mocku/pkg/common"
	"mocku/pkg/model"
	"mocku/pkg/service"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/romsar/gonertia"
	inertia "github.com/romsar/gonertia"
)

type Directive struct {
	i        *gonertia.Inertia
	services *service.Services
}

func NewDirective(i *gonertia.Inertia, services *service.Services) *Directive {
	return &Directive{
		i:        i,
		services: services,
	}
}

func (d *Directive) RegisterRoutes(e *echo.Group) {
	e.GET("/", d.DirectiveDash)
	e.GET("/students", d.Students)
	e.GET("/students/view", d.Student)
	e.POST("/students/view/submit", d.StudentPost)
	e.GET("/careers", d.Careers)
	e.POST("/careers/submit", d.Career)
	e.GET("/professors", d.Professors)
	e.GET("/professors/view", d.Professor)
	e.POST("/professors/view/submit", d.ProfessorPost)
	e.GET("/subjects", d.Subjects)
	e.GET("/subjects/view", d.Subject)
	e.POST("/subjects/view/submit", d.SubjectPost)
}

func (d *Directive) DirectiveDash(c echo.Context) error {
	careers, err := d.services.Careers.GetAll(c.Request().Context())
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Dash", gonertia.Props{
		"careers": careers,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) Students(c echo.Context) error {
	students, err := d.services.Student.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	studentDtos := make([]model.StudentsTableResponse, len(students))
	for i, student := range students {
		studentDtos[i] = model.StudentsTableResponse{
			ID:           student.ID,
			Name:         student.User.Name,
			Avatar:       student.User.Avatar,
			Email:        student.User.Email,
			Phone:        student.Phone,
			Career:       student.Career.Name,
			TotalAverage: student.TotalAverage,
		}
	}

	var dtos []interface{}
	for _, studentDto := range studentDtos {
		dtos = append(dtos, studentDto)
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Students/Home", gonertia.Props{
		"students": dtos,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) Student(c echo.Context) error {
	id := c.QueryParam("id")

	var studentDto model.StudentResponse
	var userDto model.UserResponse

	if id != "add" {
		student, err := d.services.Student.GetByID(c.Request().Context(), id)
		if err != nil {
			return nil
		}

		studentDto = model.StudentResponse{
			ID:                     id,
			Phone:                  student.Phone,
			Address:                student.Address,
			District:               student.District,
			City:                   student.City,
			PostalCode:             student.PostalCode,
			IdentityCard:           student.IdentityCard,
			BirthDate:              student.BirthDate.Format("2006-01-02"),
			CreditUnitsAccumulated: student.CreditUnitsAccumulated,
			TotalAverage:           student.TotalAverage,
		}

		userDto = model.UserResponse{
			ID:       student.User.ID,
			Name:     student.User.Name,
			Email:    student.User.Email,
			Username: student.User.Username,
			Avatar:   strings.Replace(student.User.Avatar, "./", "/", 1),
			Active:   student.User.IsActive,
		}
	}

	careers, err := d.services.Careers.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	var careerDtos []model.SelectResponse
	for _, career := range careers {
		careerDtos = append(careerDtos, model.SelectResponse{
			ID:   career.ID,
			Name: career.Name,
		})
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Students/Upsert", gonertia.Props{
		"student": studentDto,
		"user":    userDto,
		"careers": careerDtos,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) Careers(c echo.Context) error {
	careers, err := d.services.Careers.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	careerDtos := make([]model.CareerResponse, len(careers))
	for i, career := range careers {
		if career.Leader == nil {
			careerDtos[i] = model.CareerResponse{
				ID:          career.ID,
				Name:        career.Name,
				Description: career.Description,
				LeaderName:  "",
				LeaderId:    "",
			}
			continue
		}
		careerDtos[i] = model.CareerResponse{
			ID:          career.ID,
			Name:        career.Name,
			Description: career.Description,
			LeaderName:  career.Leader.User.Name,
			LeaderId:    career.Leader.User.ID,
		}
	}

	professors, err := d.services.Professor.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	professorsDto := make([]model.SelectResponse, len(professors))
	for i, professor := range professors {
		professorsDto[i] = model.SelectResponse{
			ID:   professor.ID,
			Name: professor.User.Name,
		}
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Careers", inertia.Props{
		"careers":    careerDtos,
		"professors": professorsDto,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) Professors(c echo.Context) error {
	professors, err := d.services.Professor.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	professorDtos := make([]model.ProfessorResponse, len(professors))
	for i, professor := range professors {
		professorDtos[i] = model.ProfessorResponse{
			ID:           professor.ID,
			Name:         professor.User.Name,
			Email:        professor.User.Email,
			Avatar:       strings.Replace(professor.User.Avatar, "./", "/", 1),
			IdentityCard: professor.IdentityCard,
			Phone:        professor.Phone,
		}
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Professor/Home", inertia.Props{
		"professors": professorDtos,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) Professor(c echo.Context) error {
	id := c.QueryParam("id")

	var professorDto model.ProfessorResponse
	var userDto model.UserResponse

	if id != "add" {
		professor, err := d.services.Professor.GetByID(c.Request().Context(), id)
		if err != nil {
			return nil
		}

		professorDto = model.ProfessorResponse{
			ID:           professor.ID,
			IdentityCard: professor.IdentityCard,
			Phone:        professor.Phone,
		}

		userDto = model.UserResponse{
			ID:       "",
			Name:     professor.User.Name,
			Email:    professor.User.Email,
			Username: professor.User.Username,
			Avatar:   strings.Replace(professor.User.Avatar, "./", "/", 1),
			Active:   professor.User.IsActive,
		}
	}

	// get all professors that don't have a boss, without BossID, use Boss Edge
	professors, err := d.services.Professor.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	bossesDto := make([]model.SelectResponse, len(professors))
	for i, professor := range professors {
		bossesDto[i] = model.SelectResponse{
			ID:   professor.ID,
			Name: professor.User.Name,
		}
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Professor/Upsert", inertia.Props{
		"professor": professorDto,
		"user":      userDto,
		"bosses":    bossesDto,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) Subjects(c echo.Context) error {
	subjects, err := d.services.Subject.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	subjectDtos := make([]model.SubjectResponse, len(subjects))
	for i, subj := range subjects {
		subjectDtos[i] = model.SubjectResponse{
			ID:            subj.ID,
			Name:          subj.Name,
			Description:   subj.Description,
			Code:          subj.Code,
			CreditUnits:   subj.CreditUnits,
			Semester:      subj.Semester,
			PracticeHours: subj.PracticeHours,
			TheoryHours:   subj.TheoryHours,
			LabHours:      subj.LabHours,
			TotalHours:    subj.TotalHours,
			ClassSchedule: subj.ClassSchedule,
			ProfessorId:   subj.ProfessorID,
			ProfessorName: subj.Professor.User.Name,
			Careers:       nil,
		}

		// Now we only have a single career
		subjectDtos[i].Careers = append(subjectDtos[i].Careers, model.SelectResponse{
			ID:   subj.Career.ID,
			Name: subj.Career.Name,
		})
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Subjects/Home", inertia.Props{
		"subjects": subjectDtos,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) Subject(c echo.Context) error {
	id := c.QueryParam("id")

	var subjectDto model.SubjectResponse

	if id != "add" {
		subj, err := d.services.Subject.GetByID(c.Request().Context(), id)
		if err != nil {
			return nil
		}

		subjectDto = model.SubjectResponse{
			ID:            subj.ID,
			Name:          subj.Name,
			Description:   subj.Description,
			Code:          subj.Code,
			CreditUnits:   subj.CreditUnits,
			Semester:      subj.Semester,
			PracticeHours: subj.PracticeHours,
			TheoryHours:   subj.TheoryHours,
			LabHours:      subj.LabHours,
			TotalHours:    subj.TotalHours,
			ClassSchedule: subj.ClassSchedule,
			ProfessorId:   subj.ProfessorID,
			ProfessorName: subj.Professor.User.Name,
		}

		// Now we only have a single career
		subjectDto.Careers = []model.SelectResponse{{
			ID:   subj.CareerID,
			Name: subj.Career.Name,
		}}

		// Convert prerequisites from the new structure
		var prerequisiteSubjects []model.Subject
		for _, prereq := range subj.Prerequisites {
			prerequisiteSubjects = append(prerequisiteSubjects, prereq.Prerequisite)
		}
		subjectDto.Prerequisites = model.FillSelectResponseSubject(prerequisiteSubjects)
	}

	careers, err := d.services.Careers.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}
	careerDtos := model.FillSelectResponse(careers, "ID", "Name")

	professors, err := d.services.Professor.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}
	professorDtos := make([]model.SelectResponse, len(professors))
	for i, professor := range professors {
		professorDtos[i] = model.SelectResponse{
			ID:   professor.ID,
			Name: professor.User.Name,
		}
	}

	subjects, err := d.services.Subject.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}
	subjectDtos := model.FillSelectResponse(subjects, "ID", "Name")

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Subjects/Upsert", inertia.Props{
		"subject":    subjectDto,
		"professors": professorDtos,
		"careers":    careerDtos,
		"subjects":   subjectDtos,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) StudentPost(c echo.Context) error {
	var studentRequest model.StudentRequest
	if err := c.Bind(&studentRequest); err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	handler, err := c.FormFile("avatar")
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	// Guarda el archivo si se ha subido
	filePath, err := common.UploadAvatar(studentRequest.Username, handler)
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	hashedPassword, err := common.HashPassword(studentRequest.IdentityCard)
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	user, err := d.services.Users.Create(c.Request().Context(), model.Users{
		Name:     studentRequest.Name,
		Email:    studentRequest.Email,
		Username: studentRequest.Username,
		Password: hashedPassword,
		Avatar:   filePath,
	})
	if err != nil {
		return nil
	}

	_, err = d.services.Student.Create(c.Request().Context(), model.Student{
		Phone:                  studentRequest.Phone,
		District:               studentRequest.District,
		City:                   studentRequest.City,
		Address:                studentRequest.Address,
		IdentityCard:           studentRequest.IdentityCard,
		PostalCode:             studentRequest.PostalCode,
		CreditUnitsAccumulated: studentRequest.CreditUnitsAccumulated,
		Semester:               studentRequest.Semester,
		TotalAverage:           studentRequest.TotalAverage,
		BirthDate:              studentRequest.BirthDate,
		CareerID:               studentRequest.CareerId,
		UserID:                 user.ID,
	})
	if err != nil {
		return nil
	}

	d.i.Redirect(c.Response().Writer, c.Request(), "/directive/students", 302)
	return nil
}

func (d *Directive) Career(c echo.Context) error {
	var careerRequest model.CareerRequest
	if err := c.Bind(&careerRequest); err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	if _, err := d.services.Careers.Create(c.Request().Context(), model.Careers{
		Name:        careerRequest.Name,
		Description: careerRequest.Description,
	}); err != nil {
		return nil
	}

	d.i.Redirect(c.Response().Writer, c.Request(), "/directive/careers", 302)

	return nil
}

func (d *Directive) ProfessorPost(c echo.Context) error {
	var professorRequest model.ProfessorRequest
	if err := c.Bind(&professorRequest); err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
	handler, err := c.FormFile("avatar")
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
	}

	// Guarda el archivo si se ha subido
	filePath, err := common.UploadAvatar(professorRequest.Username, handler)
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	hashedPassword, err := common.HashPassword(professorRequest.IdentityCard)
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	user, err := d.services.Users.Create(c.Request().Context(), model.Users{
		Name:     professorRequest.Name,
		Email:    professorRequest.Email,
		Username: professorRequest.Username,
		Password: hashedPassword,
		Avatar:   filePath,
	})
	if err != nil {
		return nil
	}

	// Set BossID if it exists
	var bossID *string
	if professorRequest.BossId != nil {
		bossIDValue := *professorRequest.BossId
		bossID = &bossIDValue
	}

	_, err = d.services.Professor.Create(c.Request().Context(), model.Professor{
		IdentityCard: professorRequest.IdentityCard,
		Phone:        professorRequest.Phone,
		Address:      professorRequest.Address,
		BirthDate:    professorRequest.BirthDate,
		UserID:       user.ID,
		BossID:       bossID,
	})
	if err != nil {
		return nil
	}

	d.i.Redirect(c.Response().Writer, c.Request(), "/directive/professors", 302)

	return nil
}

func (d *Directive) SubjectPost(c echo.Context) error {
	var subjectRequest model.SubjectRequest
	if err := c.Bind(&subjectRequest); err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	var classSchedule map[string][]string
	if err := common.Unmarshal(subjectRequest.ClassSchedule, &classSchedule); err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	subjectsIdsSlice := strings.Split(subjectRequest.PreqIds, ",")

	// Create the subject first
	subject, err := d.services.Subject.Create(c.Request().Context(), model.Subject{
		Name:          subjectRequest.Name,
		Description:   subjectRequest.Description,
		Code:          subjectRequest.Code,
		CreditUnits:   subjectRequest.CreditUnits,
		Semester:      subjectRequest.Semester,
		PracticeHours: subjectRequest.PracticeHours,
		TheoryHours:   subjectRequest.TheoryHours,
		LabHours:      subjectRequest.LabHours,
		TotalHours:    subjectRequest.TotalHours,
		ClassSchedule: classSchedule,
		ProfessorID:   subjectRequest.ProfessorId,
		CareerID:      subjectRequest.CareerId,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	// Now create prerequisite relationships if any
	for _, prereqID := range subjectsIdsSlice {
		// Create a prerequisite entity for each prerequisite subject
		_, err := d.services.Subject.CreatePrerequisite(c.Request().Context(), model.Prerequisite{
			SubjectID:      subject.ID,
			PrerequisiteID: prereqID,
		})
		if err != nil {
			common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}
	}

	d.i.Redirect(c.Response().Writer, c.Request(), "/directive/subjects", 302)
	return nil
}
