package schema

import (
	"fmt"
	"net/mail"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(100),
		field.String("email").MaxLen(50).NotEmpty().Unique().Validate(func(s string) error {
			if _, err := mail.ParseAddress(s); err != nil {
				return fmt.Errorf("invalid email format: %w", err)
			}
			return nil
		}),
		field.String("password").NotEmpty(),
		field.Int("class_id").Immutable(),
	}
}

func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attempts", Attempt.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("problems", Problem.Type),
		edge.From("class", Class.Type).Ref("students").Field("class_id").Required().Immutable().Unique(),
	}
}
