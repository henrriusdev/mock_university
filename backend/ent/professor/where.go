// Code generated by ent, DO NOT EDIT.

package professor

import (
	"mocku/backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldID, id))
}

// IdentityCard applies equality check predicate on the "identity_card" field. It's identical to IdentityCardEQ.
func IdentityCard(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldIdentityCard, v))
}

// BirthDate applies equality check predicate on the "birth_date" field. It's identical to BirthDateEQ.
func BirthDate(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldBirthDate, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldPhone, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldAddress, v))
}

// IdentityCardEQ applies the EQ predicate on the "identity_card" field.
func IdentityCardEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldIdentityCard, v))
}

// IdentityCardNEQ applies the NEQ predicate on the "identity_card" field.
func IdentityCardNEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldIdentityCard, v))
}

// IdentityCardIn applies the In predicate on the "identity_card" field.
func IdentityCardIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldIdentityCard, vs...))
}

// IdentityCardNotIn applies the NotIn predicate on the "identity_card" field.
func IdentityCardNotIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldIdentityCard, vs...))
}

// IdentityCardGT applies the GT predicate on the "identity_card" field.
func IdentityCardGT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldIdentityCard, v))
}

// IdentityCardGTE applies the GTE predicate on the "identity_card" field.
func IdentityCardGTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldIdentityCard, v))
}

// IdentityCardLT applies the LT predicate on the "identity_card" field.
func IdentityCardLT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldIdentityCard, v))
}

// IdentityCardLTE applies the LTE predicate on the "identity_card" field.
func IdentityCardLTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldIdentityCard, v))
}

// IdentityCardContains applies the Contains predicate on the "identity_card" field.
func IdentityCardContains(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContains(FieldIdentityCard, v))
}

// IdentityCardHasPrefix applies the HasPrefix predicate on the "identity_card" field.
func IdentityCardHasPrefix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasPrefix(FieldIdentityCard, v))
}

// IdentityCardHasSuffix applies the HasSuffix predicate on the "identity_card" field.
func IdentityCardHasSuffix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasSuffix(FieldIdentityCard, v))
}

// IdentityCardEqualFold applies the EqualFold predicate on the "identity_card" field.
func IdentityCardEqualFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEqualFold(FieldIdentityCard, v))
}

// IdentityCardContainsFold applies the ContainsFold predicate on the "identity_card" field.
func IdentityCardContainsFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContainsFold(FieldIdentityCard, v))
}

// BirthDateEQ applies the EQ predicate on the "birth_date" field.
func BirthDateEQ(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldBirthDate, v))
}

// BirthDateNEQ applies the NEQ predicate on the "birth_date" field.
func BirthDateNEQ(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldBirthDate, v))
}

// BirthDateIn applies the In predicate on the "birth_date" field.
func BirthDateIn(vs ...time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldBirthDate, vs...))
}

// BirthDateNotIn applies the NotIn predicate on the "birth_date" field.
func BirthDateNotIn(vs ...time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldBirthDate, vs...))
}

// BirthDateGT applies the GT predicate on the "birth_date" field.
func BirthDateGT(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldBirthDate, v))
}

// BirthDateGTE applies the GTE predicate on the "birth_date" field.
func BirthDateGTE(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldBirthDate, v))
}

// BirthDateLT applies the LT predicate on the "birth_date" field.
func BirthDateLT(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldBirthDate, v))
}

// BirthDateLTE applies the LTE predicate on the "birth_date" field.
func BirthDateLTE(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldBirthDate, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContainsFold(FieldPhone, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContainsFold(FieldAddress, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.Users) predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBoss applies the HasEdge predicate on the "boss" edge.
func HasBoss() predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BossTable, BossColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBossWith applies the HasEdge predicate on the "boss" edge with a given conditions (other predicates).
func HasBossWith(preds ...predicate.Professor) predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := newBossStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubordinates applies the HasEdge predicate on the "subordinates" edge.
func HasSubordinates() predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubordinatesTable, SubordinatesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubordinatesWith applies the HasEdge predicate on the "subordinates" edge with a given conditions (other predicates).
func HasSubordinatesWith(preds ...predicate.Professor) predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := newSubordinatesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubjects applies the HasEdge predicate on the "subjects" edge.
func HasSubjects() predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, SubjectsTable, SubjectsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectsWith applies the HasEdge predicate on the "subjects" edge with a given conditions (other predicates).
func HasSubjectsWith(preds ...predicate.Subject) predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := newSubjectsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCareers applies the HasEdge predicate on the "careers" edge.
func HasCareers() predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, CareersTable, CareersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCareersWith applies the HasEdge predicate on the "careers" edge with a given conditions (other predicates).
func HasCareersWith(preds ...predicate.Careers) predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := newCareersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Professor) predicate.Professor {
	return predicate.Professor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Professor) predicate.Professor {
	return predicate.Professor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Professor) predicate.Professor {
	return predicate.Professor(sql.NotPredicates(p))
}
