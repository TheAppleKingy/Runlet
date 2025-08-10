package schema

import (
	"fmt"
	"net/mail"

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
		field.String("email").MaxLen(50).NotEmpty().Unique().Validate(func(s string) error {
			if _, err := mail.ParseAddress(s); err != nil {
				return fmt.Errorf("invalid email format: %w", err)
			}
			return nil
		}),
		field.String("password").NotEmpty(),
	}
}

func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("classes", Class.Type),
		edge.To("courses", Course.Type),
	}
}
