package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Course struct {
	ent.Schema
}

func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique(),
		field.String("description").Unique(),
	}
}

func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("problems", Problem.Type),
		edge.From("classes", Class.Type).Ref("courses"),
		edge.From("teachers", Teacher.Type).Ref("courses"),
	}
}
