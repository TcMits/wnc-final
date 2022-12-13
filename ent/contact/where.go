// Code generated by ent, DO NOT EDIT.

package contact

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/TcMits/wnc-final/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerID), v))
	})
}

// AccountNumber applies equality check predicate on the "account_number" field. It's identical to AccountNumberEQ.
func AccountNumber(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountNumber), v))
	})
}

// SuggestName applies equality check predicate on the "suggest_name" field. It's identical to SuggestNameEQ.
func SuggestName(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSuggestName), v))
	})
}

// BankName applies equality check predicate on the "bank_name" field. It's identical to BankNameEQ.
func BankName(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBankName), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerID), v))
	})
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v uuid.UUID) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOwnerID), v))
	})
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...uuid.UUID) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOwnerID), v...))
	})
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...uuid.UUID) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOwnerID), v...))
	})
}

// AccountNumberEQ applies the EQ predicate on the "account_number" field.
func AccountNumberEQ(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberNEQ applies the NEQ predicate on the "account_number" field.
func AccountNumberNEQ(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberIn applies the In predicate on the "account_number" field.
func AccountNumberIn(vs ...string) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccountNumber), v...))
	})
}

// AccountNumberNotIn applies the NotIn predicate on the "account_number" field.
func AccountNumberNotIn(vs ...string) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccountNumber), v...))
	})
}

// AccountNumberGT applies the GT predicate on the "account_number" field.
func AccountNumberGT(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberGTE applies the GTE predicate on the "account_number" field.
func AccountNumberGTE(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberLT applies the LT predicate on the "account_number" field.
func AccountNumberLT(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberLTE applies the LTE predicate on the "account_number" field.
func AccountNumberLTE(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberContains applies the Contains predicate on the "account_number" field.
func AccountNumberContains(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberHasPrefix applies the HasPrefix predicate on the "account_number" field.
func AccountNumberHasPrefix(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberHasSuffix applies the HasSuffix predicate on the "account_number" field.
func AccountNumberHasSuffix(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberEqualFold applies the EqualFold predicate on the "account_number" field.
func AccountNumberEqualFold(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberContainsFold applies the ContainsFold predicate on the "account_number" field.
func AccountNumberContainsFold(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccountNumber), v))
	})
}

// SuggestNameEQ applies the EQ predicate on the "suggest_name" field.
func SuggestNameEQ(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSuggestName), v))
	})
}

// SuggestNameNEQ applies the NEQ predicate on the "suggest_name" field.
func SuggestNameNEQ(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSuggestName), v))
	})
}

// SuggestNameIn applies the In predicate on the "suggest_name" field.
func SuggestNameIn(vs ...string) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSuggestName), v...))
	})
}

// SuggestNameNotIn applies the NotIn predicate on the "suggest_name" field.
func SuggestNameNotIn(vs ...string) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSuggestName), v...))
	})
}

// SuggestNameGT applies the GT predicate on the "suggest_name" field.
func SuggestNameGT(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSuggestName), v))
	})
}

// SuggestNameGTE applies the GTE predicate on the "suggest_name" field.
func SuggestNameGTE(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSuggestName), v))
	})
}

// SuggestNameLT applies the LT predicate on the "suggest_name" field.
func SuggestNameLT(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSuggestName), v))
	})
}

// SuggestNameLTE applies the LTE predicate on the "suggest_name" field.
func SuggestNameLTE(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSuggestName), v))
	})
}

// SuggestNameContains applies the Contains predicate on the "suggest_name" field.
func SuggestNameContains(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSuggestName), v))
	})
}

// SuggestNameHasPrefix applies the HasPrefix predicate on the "suggest_name" field.
func SuggestNameHasPrefix(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSuggestName), v))
	})
}

// SuggestNameHasSuffix applies the HasSuffix predicate on the "suggest_name" field.
func SuggestNameHasSuffix(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSuggestName), v))
	})
}

// SuggestNameEqualFold applies the EqualFold predicate on the "suggest_name" field.
func SuggestNameEqualFold(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSuggestName), v))
	})
}

// SuggestNameContainsFold applies the ContainsFold predicate on the "suggest_name" field.
func SuggestNameContainsFold(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSuggestName), v))
	})
}

// BankNameEQ applies the EQ predicate on the "bank_name" field.
func BankNameEQ(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBankName), v))
	})
}

// BankNameNEQ applies the NEQ predicate on the "bank_name" field.
func BankNameNEQ(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBankName), v))
	})
}

// BankNameIn applies the In predicate on the "bank_name" field.
func BankNameIn(vs ...string) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBankName), v...))
	})
}

// BankNameNotIn applies the NotIn predicate on the "bank_name" field.
func BankNameNotIn(vs ...string) predicate.Contact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBankName), v...))
	})
}

// BankNameGT applies the GT predicate on the "bank_name" field.
func BankNameGT(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBankName), v))
	})
}

// BankNameGTE applies the GTE predicate on the "bank_name" field.
func BankNameGTE(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBankName), v))
	})
}

// BankNameLT applies the LT predicate on the "bank_name" field.
func BankNameLT(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBankName), v))
	})
}

// BankNameLTE applies the LTE predicate on the "bank_name" field.
func BankNameLTE(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBankName), v))
	})
}

// BankNameContains applies the Contains predicate on the "bank_name" field.
func BankNameContains(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBankName), v))
	})
}

// BankNameHasPrefix applies the HasPrefix predicate on the "bank_name" field.
func BankNameHasPrefix(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBankName), v))
	})
}

// BankNameHasSuffix applies the HasSuffix predicate on the "bank_name" field.
func BankNameHasSuffix(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBankName), v))
	})
}

// BankNameEqualFold applies the EqualFold predicate on the "bank_name" field.
func BankNameEqualFold(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBankName), v))
	})
}

// BankNameContainsFold applies the ContainsFold predicate on the "bank_name" field.
func BankNameContainsFold(v string) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBankName), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Customer) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Contact) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Contact) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Contact) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		p(s.Not())
	})
}
