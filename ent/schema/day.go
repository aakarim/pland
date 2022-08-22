package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Day holds the schema definition for the Day entity.
type Day struct {
	ent.Schema
}

// Fields of the Day.
func (Day) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Annotations(
			entgql.OrderField("CREATED_AT"),
		).Default(time.Now),
		field.Time("date").Annotations(
			entgql.OrderField("DATE"),
		),
		field.Text("txt"),
	}
}

// Edges of the Day.
func (Day) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).Ref("days"),
	}
}
