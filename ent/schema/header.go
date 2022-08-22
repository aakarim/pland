package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Header holds the schema definition for the Header entity.
type Header struct {
	ent.Schema
}

// Fields of the Header.
func (Header) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Annotations(
			entgql.OrderField("CREATED_AT"),
		).Default(time.Now),
		field.Text("txt"),
	}
}

// Edges of the Header.
func (Header) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).Ref("header").Unique(),
	}
}
