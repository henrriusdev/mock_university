package handlers

import (
	"fmt"
	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"net/http"
	"strconv"

	"mocku/backend/ent/users"
	"mocku/backend/utils"

	inertia "github.com/romsar/gonertia"
)

func (h *Handler) HomeHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Home/Index", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) LoginHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Auth/Login", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) LoginPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		var formData LoginDto
		// get formData from request instead of json
		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err)
			return
		}

		formData.Email = r.FormValue("email")
		formData.Password = r.FormValue("password")
		user, err := h.DB.Users.Query().Where(users.EmailEQ(formData.Email)).WithRole().First(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}
		if !utils.CheckPassword(user.Password, formData.Password) {
			err = i.Render(w, r, "Auth/Login", inertia.Props{
				"careers": careers,
				"error":   "Credenciales incorrectas",
			})
			if err != nil {
				HandleServerErr(i, err)
				return
			}

			return
		}

		var view string
		role := user.Edges.Role.ID
		switch role {
		case 1:
			view = "/directive"
		case 2:
			view = "/payments"
		case 3:
			view = "/control"
		case 4, 5:
			view = "/professor"
		case 6:
			view = "/student"
		}

		println(view)

		i.Redirect(w, r, view, 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) DirectiveDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		fmt.Println(err)
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Directive/Dash", inertia.Props{
			"careers": careers,
		})
		fmt.Println(err)
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) PaymentsDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Payment/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) ControlDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Control/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) ProfessorDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Professor/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) StudentDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Student/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		config := h.DB.Configuration.Query().WithCycle().Where(configuration.HasCycleWith(cycle.Name("2024-2"))).OnlyX(r.Context())
		if config == nil {
			HandleServerErr(i, fmt.Errorf("config not found")).ServeHTTP(w, r)
			return
		}

		err := i.Render(w, r, "Directive/Settings/Home", inertia.Props{
			"notesNumber":   config.NumberNotes,
			"paymentNumber": config.NumberFees,
			"startRegSubj":  utils.FormatDate(config.StartRegistrationSubjects),
			"endRegSubj":    utils.FormatDate(config.EndRegistrationSubjects),
			"cycleStart":    utils.FormatDate(config.Edges.Cycle.StartDate),
			"cycleEnd":      utils.FormatDate(config.Edges.Cycle.EndDate),
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsProfileHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Settings/Profile", nil)
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsProfilePostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		err := i.Render(w, r, "Settings/Profile", nil)
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsNotesPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		notesNumber, err := strconv.Atoi(r.FormValue("notes"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		_, err = h.DB.Configuration.Update().Where(configuration.ID(1)).SetNumberNotes(notesNumber).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		i.Redirect(w, r, "/settings", 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsDatesPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		startRegistrationSubjects, err := utils.ParseDate(r.FormValue("start_registration_subjects"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		endRegistrationSubjects, err := utils.ParseDate(r.FormValue("end_registration_subjects"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		cycleStart, err := utils.ParseDate(r.FormValue("cycle_start"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		cycleEnd, err := utils.ParseDate(r.FormValue("cycle_end"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		_, err = h.DB.Cycle.Update().Where(cycle.ActiveEQ(true)).SetStartDate(cycleStart).SetEndDate(cycleEnd).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		_, err = h.DB.Configuration.Update().Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).SetStartRegistrationSubjects(startRegistrationSubjects).SetEndRegistrationSubjects(endRegistrationSubjects).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		i.Redirect(w, r, "/settings", 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsPaymentsPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		numberFees, err := strconv.Atoi(r.FormValue("payments"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		_, err = h.DB.Configuration.Update().Where(configuration.ID(1)).SetNumberFees(numberFees).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		i.Redirect(w, r, "/settings", 302)
	}

	return http.HandlerFunc(fn)
}
