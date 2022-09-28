package clevercloud

import (
	"context"
	"fmt"
	"reflect"

	"github.com/fatih/structtag"
	"github.com/gobeam/stringy"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/quals"
	"go.clever-cloud.dev/sdk/models"
	"go.clever-cloud.dev/sdk/v2/organisation"
)

// Compute column list from a struct
func extractColumns(ctx context.Context, i interface{}) []*plugin.Column {
	log := plugin.Logger(ctx)
	columns := []*plugin.Column{}

	t := reflect.TypeOf(i)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			panic(err)
		}

		log.Debug(fmt.Sprintf("Field %+v", field))
		col := &plugin.Column{
			Name: field.Name,
			Type: typeMapper(field.Type),
		}

		if jsonTag, err := tags.Get("json"); err == nil {
			name := stringy.New(jsonTag.Name)
			col.Name = name.SnakeCase("?", "").ToLower()
		}

		if defaultTag, err := tags.Get("default"); err == nil {
			col.Default = defaultTag.Name
		}

		if descTag, err := tags.Get("description"); err == nil {
			col.Description = descTag.Name
		}

		if col.Type != proto.ColumnType_UNKNOWN {
			columns = append(columns, col)
		} else {
			log.Error("unsupported struct field", "type", field.Type.Name())
		}
	}

	return columns
}

// map a go type to a steampipe protobuf one
func typeMapper(t reflect.Type) proto.ColumnType {
	switch t.Name() {
	case "string":
		return proto.ColumnType_STRING
	case "bool":
		return proto.ColumnType_BOOL
	case "int32", "int64":
		return proto.ColumnType_INT
	default:
		return proto.ColumnType_UNKNOWN
	}
}

func Map[T, U any](items []T, fn func(T) U) []U {
	r := make([]U, len(items))

	for i, item := range items {
		r[i] = fn(item)
	}

	return r
}

func Foreach[T any](items []T, fn func(T)) {
	for _, item := range items {
		fn(item)
	}
}

// GetOrgsList return a list od organisation ID from sql clauses or from API
func GetOrgsList(ctx context.Context, d *plugin.QueryData) ([]string, error) {
	cc := getClient()

	if d.Quals["owner_id"] != nil {
		return Map(d.Quals["owner_id"].Quals, func(q *quals.Qual) string {
			return q.Value.GetStringValue()
		}), nil
	}

	orgRes := organisation.List(ctx, cc)
	if orgRes.HasError() {
		return nil, orgRes.Error()
	}

	return Map(*orgRes.Payload(), func(org models.OrganisationView) string {
		return org.ID
	}), nil
}
