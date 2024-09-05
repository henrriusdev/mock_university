package repos

import (
	inertia "github.com/romsar/gonertia"
	"mocku/backend/common"
	"mocku/backend/ent"
	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"net/http"
	"time"
)

func (r *Repo) UpdateNumberNotes(notesNumber int, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Configuration.Update().
		Where(configuration.HasCycleWith(cycle.Active(true))).
		SetNumberNotes(notesNumber).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating configuration: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) UpdateDates(startSubjects, endSubjects, cycleStart, cycleEnd time.Time, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Cycle.Update().
		Where(cycle.ActiveEQ(true)).
		SetStartDate(cycleStart).
		SetEndDate(cycleEnd).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating cycle: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	_, err = r.DB.Configuration.Update().
		Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).
		SetStartRegistrationSubjects(startSubjects).
		SetEndRegistrationSubjects(endSubjects).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating configuration: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) UpdateNumberFees(feesNumber int, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Configuration.Update().
		Where(configuration.HasCycleWith(cycle.Active(true))).
		SetNumberFees(feesNumber).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating configuration: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) GetConfiguration(i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Configuration, error) {
	config, err := r.DB.Configuration.Query().
		Where(configuration.HasCycleWith(cycle.Active(true))).
		WithCycle().
		First(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying configuration: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return config, nil
}

func (r *Repo) UpdateNotesPercentages(percentages []float64, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Configuration.Update().
		Where(configuration.HasCycleWith(cycle.Active(true))).
		SetNotesPercentages(percentages).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating configuration: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) UpdateFeeDates(payments []time.Time, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Configuration.Update().
		Where(configuration.HasCycleWith(cycle.Active(true))).
		SetFeeDates(payments).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating configuration: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) GetCurrentCycle(i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Cycle, error) {
	current, err := r.DB.Cycle.Query().
		Where(cycle.ActiveEQ(true)).
		First(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying cycle: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return current, nil
}

func (r *Repo) InactivateCycle(i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Cycle.Update().
		Where(cycle.ActiveEQ(true)).
		SetActive(false).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error inactivating cycle: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) NewCycle(name string, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Cycle, error) {
	newCycle, err := r.DB.Cycle.Create().
		SetActive(true).
		SetName(name).
		SetStartDate(time.Now()).
		SetEndDate(time.Now()).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error creating cycle: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return newCycle, nil
}

func (r *Repo) NewConfiguration(currentCycle *ent.Cycle, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Configuration.Create().
		SetNumberNotes(0).
		SetNumberFees(0).
		SetStartRegistrationSubjects(time.Now()).
		SetEndRegistrationSubjects(time.Now()).
		SetNotesPercentages([]float64{}).
		SetFeeDates([]time.Time{}).
		SetCycle(currentCycle).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error creating configuration: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}
