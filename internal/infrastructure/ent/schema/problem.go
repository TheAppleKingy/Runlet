package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Problem struct {
	ent.Schema
}

func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(70),
		field.String("description"),
		field.Int("course_id"),
	}
}

func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).Ref("problems").Field("course_id").Unique().Required(),
		edge.To("attempts", Attempt.Type),
		edge.From("students", Student.Type).Ref("problems"),
	}
}
