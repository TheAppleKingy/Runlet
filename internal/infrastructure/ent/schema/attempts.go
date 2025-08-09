package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Attempt struct {
	ent.Schema
}

func (Attempt) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("amount").Default(0),
		field.Bool("done").Default(false),
		field.Int("student_id"),
		field.Int("problem_id"),
	}
}

func (Attempt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", Student.Type).Ref("attempts").Field("student_id").Unique().Required(),
		edge.From("problem", Problem.Type).Ref("attempts").Field("problem_id").Unique().Required(),
	}
}
