package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Teacher struct {
	ent.Schema
}

func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(100),
	}
}

func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("classes", Class.Type),
		edge.To("courses", Course.Type),
	}
}
