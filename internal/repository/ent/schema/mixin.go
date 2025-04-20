package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// IDMixin implements the ent.Mixin for sharing the ID field
// with all schemas that embed it.
type IDMixin struct {
	mixin.Schema
}

// Fields of the IDMixin.
func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive().
			Immutable().
			Comment("The ID of the entity").
			StructTag(`json:"id,omitempty"`),
	}
}

// TimeMixin implements the ent.Mixin for sharing
// created_at and updated_at fields with all schemas that embed it.
type TimeMixin struct {
	mixin.Schema
}

// Fields of the TimeMixin.
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("The time when the entity was created"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("The time when the entity was last updated"),
	}
}

// SoftDeleteMixin implements the ent.Mixin for sharing
// soft delete functionality with all schemas that embed it.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Optional().
			Nillable().
			Comment("The time when the entity was soft deleted"),
	}
}
