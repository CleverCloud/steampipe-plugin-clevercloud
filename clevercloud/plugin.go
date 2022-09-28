package clevercloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	log := plugin.Logger(ctx)

	log.Trace("Plugin()")

	p := &plugin.Plugin{
		Name:             "steampipe-plugin-clevercloud",
		DefaultTransform: transform.From(camelCase),
		TableMap: map[string]*plugin.Table{
			"cc_app": AppTable(ctx),
			"cc_org": OrgTable(ctx),
		},
	}

	return p
}

func camelCase(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	propertyPath := strcase.ToCamel(d.ColumnName)
	if propertyPath == "" {
		return nil, fmt.Errorf("'FieldValue' requires a string parameter containing property path but received %v", d.Param)
	}

	propertyPath = strings.ReplaceAll(propertyPath, "Id", "ID")

	d.Param = propertyPath

	return transform.FieldValue(ctx, d)
}
