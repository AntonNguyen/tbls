package config

import (
	"database/sql"
	"testing"

	"github.com/k1LoW/tbls/schema"
)

func TestRequireTableComment(t *testing.T) {
	r := RequireTableComment{
		Enabled: true,
	}
	s := newTestSchema()
	warns := r.Check(s)
	if len(warns) != 1 {
		t.Errorf("actual %v\nwant %v", len(warns), 1)
	}
}

func TestRequireTableCommentWithExclude(t *testing.T) {
	r := RequireTableComment{
		Enabled: true,
		Exclude: []string{"a"},
	}
	s := newTestSchema()
	warns := r.Check(s)
	if len(warns) != 0 {
		t.Errorf("actual %v\nwant %v", len(warns), 1)
	}
}

func TestRequireColumnComment(t *testing.T) {
	r := RequireColumnComment{
		Enabled: true,
	}
	s := newTestSchema()
	warns := r.Check(s)
	if len(warns) != 1 {
		t.Errorf("actual %v\nwant %v", len(warns), 1)
	}
}

func TestRequireColumnCommentWithExclude(t *testing.T) {
	r := RequireColumnComment{
		Enabled: true,
		Exclude: []string{"b"},
	}
	s := newTestSchema()
	warns := r.Check(s)
	if len(warns) != 0 {
		t.Errorf("actual %v\nwant %v", len(warns), 1)
	}
}

func TestUnrelatedTable(t *testing.T) {
	r := UnrelatedTable{
		Enabled: true,
	}
	s := newTestSchema()
	warns := r.Check(s)
	if len(warns) != 1 {
		t.Errorf("actual %v\nwant %v", len(warns), 1)
	}
}

func TestUnrelatedTableWithExclude(t *testing.T) {
	r := UnrelatedTable{
		Enabled: true,
		Exclude: []string{"c"},
	}
	s := newTestSchema()
	warns := r.Check(s)
	if len(warns) != 0 {
		t.Errorf("actual %v\nwant %v", len(warns), 1)
	}
}

func TestColumnCount(t *testing.T) {
	r := ColumnCount{
		Enabled: true,
		Max:     3,
	}
	s := newTestSchema()
	warns := r.Check(s)
	if len(warns) != 1 {
		t.Errorf("actual %v\nwant %v", len(warns), 1)
	}
}

func TestColumnCountWithExclude(t *testing.T) {
	r := ColumnCount{
		Enabled: true,
		Max:     3,
		Exclude: []string{"c"},
	}
	s := newTestSchema()
	warns := r.Check(s)
	if len(warns) != 0 {
		t.Errorf("actual %v\nwant %v", len(warns), 1)
	}
}

func newTestSchema() *schema.Schema {
	ca := &schema.Column{
		Name:     "a",
		Type:     "bigint(20)",
		Comment:  "column a",
		Nullable: false,
	}
	cb := &schema.Column{
		Name:     "b",
		Type:     "text",
		Comment:  "", // empty comment
		Nullable: true,
	}

	ta := &schema.Table{
		Name:    "a",
		Type:    "BASE TABLE",
		Comment: "", // empty comment
		Columns: []*schema.Column{
			ca,
			&schema.Column{
				Name:     "a2",
				Type:     "datetime",
				Comment:  "column a2",
				Nullable: false,
				Default: sql.NullString{
					String: "CURRENT_TIMESTAMP",
					Valid:  true,
				},
			},
		},
	}
	tb := &schema.Table{
		Name:    "b",
		Type:    "BASE TABLE",
		Comment: "table b",
		Columns: []*schema.Column{
			cb,
			&schema.Column{
				Name:     "b2",
				Comment:  "column b2",
				Type:     "text",
				Nullable: true,
			},
		},
	}
	tc := &schema.Table{
		Name:    "c",
		Type:    "BASE TABLE",
		Comment: "table c",
		Columns: []*schema.Column{
			&schema.Column{
				Name:     "c1",
				Type:     "text",
				Comment:  "column c1",
				Nullable: false,
			},
			&schema.Column{
				Name:     "c2",
				Type:     "text",
				Comment:  "column c2",
				Nullable: false,
			},
			&schema.Column{
				Name:     "c3",
				Type:     "text",
				Comment:  "column c3",
				Nullable: false,
			},
			&schema.Column{
				Name:     "c4",
				Type:     "text",
				Comment:  "column c4",
				Nullable: false,
			},
		},
	}
	r := &schema.Relation{
		Table:         ta,
		Columns:       []*schema.Column{ca},
		ParentTable:   tb,
		ParentColumns: []*schema.Column{cb},
	}
	ca.ParentRelations = []*schema.Relation{r}
	cb.ChildRelations = []*schema.Relation{r}

	s := &schema.Schema{
		Name: "testschema",
		Tables: []*schema.Table{
			ta,
			tb,
			tc,
		},
		Relations: []*schema.Relation{
			r,
		},
		Driver: &schema.Driver{
			Name:            "testdriver",
			DatabaseVersion: "1.0.0",
		},
	}
	return s
}
