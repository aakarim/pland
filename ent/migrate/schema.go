// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArbitrarySectionsColumns holds the columns for the "arbitrary_sections" table.
	ArbitrarySectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "token", Type: field.TypeString},
		{Name: "txt", Type: field.TypeString, Size: 2147483647},
	}
	// ArbitrarySectionsTable holds the schema information for the "arbitrary_sections" table.
	ArbitrarySectionsTable = &schema.Table{
		Name:       "arbitrary_sections",
		Columns:    ArbitrarySectionsColumns,
		PrimaryKey: []*schema.Column{ArbitrarySectionsColumns[0]},
	}
	// DaysColumns holds the columns for the "days" table.
	DaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "date", Type: field.TypeTime},
		{Name: "txt", Type: field.TypeString, Size: 2147483647},
	}
	// DaysTable holds the schema information for the "days" table.
	DaysTable = &schema.Table{
		Name:       "days",
		Columns:    DaysColumns,
		PrimaryKey: []*schema.Column{DaysColumns[0]},
	}
	// HeadersColumns holds the columns for the "headers" table.
	HeadersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "txt", Type: field.TypeString, Size: 2147483647},
		{Name: "plan_header", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// HeadersTable holds the schema information for the "headers" table.
	HeadersTable = &schema.Table{
		Name:       "headers",
		Columns:    HeadersColumns,
		PrimaryKey: []*schema.Column{HeadersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "headers_plans_header",
				Columns:    []*schema.Column{HeadersColumns[3]},
				RefColumns: []*schema.Column{PlansColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PlansColumns holds the columns for the "plans" table.
	PlansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "has_conflict", Type: field.TypeBool, Default: false},
		{Name: "digest", Type: field.TypeString},
		{Name: "txt", Type: field.TypeString, Size: 2147483647},
		{Name: "plan_next", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "user_plans", Type: field.TypeInt, Nullable: true},
	}
	// PlansTable holds the schema information for the "plans" table.
	PlansTable = &schema.Table{
		Name:       "plans",
		Columns:    PlansColumns,
		PrimaryKey: []*schema.Column{PlansColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "plans_plans_next",
				Columns:    []*schema.Column{PlansColumns[5]},
				RefColumns: []*schema.Column{PlansColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "plans_charm_user_plans",
				Columns:    []*schema.Column{PlansColumns[6]},
				RefColumns: []*schema.Column{CharmUserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CharmUserColumns holds the columns for the "charm_user" table.
	CharmUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "charm_id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "bio", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
	}
	// CharmUserTable holds the schema information for the "charm_user" table.
	CharmUserTable = &schema.Table{
		Name:       "charm_user",
		Columns:    CharmUserColumns,
		PrimaryKey: []*schema.Column{CharmUserColumns[0]},
	}
	// PlanDaysColumns holds the columns for the "plan_days" table.
	PlanDaysColumns = []*schema.Column{
		{Name: "plan_id", Type: field.TypeInt},
		{Name: "day_id", Type: field.TypeInt},
	}
	// PlanDaysTable holds the schema information for the "plan_days" table.
	PlanDaysTable = &schema.Table{
		Name:       "plan_days",
		Columns:    PlanDaysColumns,
		PrimaryKey: []*schema.Column{PlanDaysColumns[0], PlanDaysColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "plan_days_plan_id",
				Columns:    []*schema.Column{PlanDaysColumns[0]},
				RefColumns: []*schema.Column{PlansColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "plan_days_day_id",
				Columns:    []*schema.Column{PlanDaysColumns[1]},
				RefColumns: []*schema.Column{DaysColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PlanArbitrarySectionsColumns holds the columns for the "plan_arbitrarySections" table.
	PlanArbitrarySectionsColumns = []*schema.Column{
		{Name: "plan_id", Type: field.TypeInt},
		{Name: "arbitrary_section_id", Type: field.TypeInt},
	}
	// PlanArbitrarySectionsTable holds the schema information for the "plan_arbitrarySections" table.
	PlanArbitrarySectionsTable = &schema.Table{
		Name:       "plan_arbitrarySections",
		Columns:    PlanArbitrarySectionsColumns,
		PrimaryKey: []*schema.Column{PlanArbitrarySectionsColumns[0], PlanArbitrarySectionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "plan_arbitrarySections_plan_id",
				Columns:    []*schema.Column{PlanArbitrarySectionsColumns[0]},
				RefColumns: []*schema.Column{PlansColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "plan_arbitrarySections_arbitrary_section_id",
				Columns:    []*schema.Column{PlanArbitrarySectionsColumns[1]},
				RefColumns: []*schema.Column{ArbitrarySectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArbitrarySectionsTable,
		DaysTable,
		HeadersTable,
		PlansTable,
		CharmUserTable,
		PlanDaysTable,
		PlanArbitrarySectionsTable,
	}
)

func init() {
	HeadersTable.ForeignKeys[0].RefTable = PlansTable
	PlansTable.ForeignKeys[0].RefTable = PlansTable
	PlansTable.ForeignKeys[1].RefTable = CharmUserTable
	CharmUserTable.Annotation = &entsql.Annotation{
		Table: "charm_user",
	}
	PlanDaysTable.ForeignKeys[0].RefTable = PlansTable
	PlanDaysTable.ForeignKeys[1].RefTable = DaysTable
	PlanArbitrarySectionsTable.ForeignKeys[0].RefTable = PlansTable
	PlanArbitrarySectionsTable.ForeignKeys[1].RefTable = ArbitrarySectionsTable
}
