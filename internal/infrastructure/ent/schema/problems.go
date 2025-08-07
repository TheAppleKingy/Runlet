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
		field.String("title"),
		field.String("description"),
		field.Int("course_id"),
	}
}

func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).Ref("problems").Field("course_id").Unique().Required(),
	}
}
