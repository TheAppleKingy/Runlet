package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		edge.To("students", Student.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("teachers", Teacher.Type),
		edge.To("courses", Course.Type),
	}
}
