// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeBool},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// PreferencesColumns holds the columns for the "preferences" table.
	PreferencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "free_weekends", Type: field.TypeBool, Default: false},
		{Name: "weekly_frequency", Type: field.TypeInt, Nullable: true},
		{Name: "user_preference", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PreferencesTable holds the schema information for the "preferences" table.
	PreferencesTable = &schema.Table{
		Name:       "preferences",
		Columns:    PreferencesColumns,
		PrimaryKey: []*schema.Column{PreferencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "preferences_users_preference",
				Columns:    []*schema.Column{PreferencesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeBool},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email_address", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "premium", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_uuid",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[4]},
			},
		},
	}
	// UserSkillsColumns holds the columns for the "user_skills" table.
	UserSkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "level", Type: field.TypeEnum, Enums: []string{"beginner", "intermediate", "advanced"}},
		{Name: "progress", Type: field.TypeInt, Default: 0},
		{Name: "duration", Type: field.TypeInt, Default: 0},
		{Name: "skill_userskills", Type: field.TypeInt, Nullable: true},
		{Name: "user_skills", Type: field.TypeInt, Nullable: true},
	}
	// UserSkillsTable holds the schema information for the "user_skills" table.
	UserSkillsTable = &schema.Table{
		Name:       "user_skills",
		Columns:    UserSkillsColumns,
		PrimaryKey: []*schema.Column{UserSkillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_skills_skills_userskills",
				Columns:    []*schema.Column{UserSkillsColumns[4]},
				RefColumns: []*schema.Column{SkillsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_skills_users_skills",
				Columns:    []*schema.Column{UserSkillsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategorySkillsColumns holds the columns for the "category_skills" table.
	CategorySkillsColumns = []*schema.Column{
		{Name: "category_id", Type: field.TypeInt},
		{Name: "skill_id", Type: field.TypeInt},
	}
	// CategorySkillsTable holds the schema information for the "category_skills" table.
	CategorySkillsTable = &schema.Table{
		Name:       "category_skills",
		Columns:    CategorySkillsColumns,
		PrimaryKey: []*schema.Column{CategorySkillsColumns[0], CategorySkillsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "category_skills_category_id",
				Columns:    []*schema.Column{CategorySkillsColumns[0]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "category_skills_skill_id",
				Columns:    []*schema.Column{CategorySkillsColumns[1]},
				RefColumns: []*schema.Column{SkillsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		PreferencesTable,
		SkillsTable,
		UsersTable,
		UserSkillsTable,
		CategorySkillsTable,
	}
)

func init() {
	PreferencesTable.ForeignKeys[0].RefTable = UsersTable
	UserSkillsTable.ForeignKeys[0].RefTable = SkillsTable
	UserSkillsTable.ForeignKeys[1].RefTable = UsersTable
	CategorySkillsTable.ForeignKeys[0].RefTable = CategoriesTable
	CategorySkillsTable.ForeignKeys[1].RefTable = SkillsTable
}
