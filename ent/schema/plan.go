package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Plan holds the schema definition for the Plan entity.
type Plan struct {
	ent.Schema
}

func (Plan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.Mutations(),
	}
}

// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date").Immutable().Annotations(
			entgql.OrderField("DATE"),
		),
		field.Time("created_at").Immutable().Annotations(
			entgql.OrderField("CREATED_AT"),
		),
		field.Time("timestamp").Immutable().Annotations(
			entgql.OrderField("TIMESTAMP"),
		),
		field.String("digest"),
		field.Text("txt"),
	}
}

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("plans").Unique(),
	}
}
