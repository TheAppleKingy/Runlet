package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(100),
		field.Int("class_id").Immutable(),
	}
}

func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attempts", Attempt.Type),
		edge.To("problems", Problem.Type),
		edge.From("class", Class.Type).Ref("students").Field("class_id").Required().Immutable().Unique(),
	}
}
