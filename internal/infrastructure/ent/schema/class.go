package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Class struct {
	ent.Schema
}

func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").Match(regexp.MustCompile(`^\d{6}$`)).Immutable().Unique(),
	}
}

func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("students", Student.Type),
		edge.From("teachers", Teacher.Type).Ref("classes"),
		edge.To("courses", Course.Type),
	}
}
