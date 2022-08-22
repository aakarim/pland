package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArbitrarySection holds the schema definition for the ArbitrarySection entity.
type ArbitrarySection struct {
	ent.Schema
}

// Fields of the ArbitrarySection.
func (ArbitrarySection) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Annotations(
			entgql.OrderField("CREATED_AT"),
		).Default(time.Now),
		field.String("token"),
		field.Text("txt"),
	}
}

// Edges of the ArbitrarySection.
func (ArbitrarySection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).Ref("arbitrarySections"),
	}
}
