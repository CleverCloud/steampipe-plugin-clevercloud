package clevercloud

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"go.clever-cloud.dev/sdk/models"
	"go.clever-cloud.dev/sdk/v2/organisation"
)

func OrgTable(ctx context.Context) *plugin.Table {
	log := plugin.Logger(ctx)
	log.Trace("OrgTable()")

	appTable := &plugin.Table{
		Name:        "cc_org",
		Description: "CleverCloud organisations",
		Columns:     extractColumns(ctx, models.OrganisationView{}),
		List:        &plugin.ListConfig{Hydrate: listOrgs},
		Get:         &plugin.GetConfig{Hydrate: getOrg, KeyColumns: plugin.SingleColumn("id")},
	}

	return appTable
}

func listOrgs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log := plugin.Logger(ctx)
	cc := getClient()

	log.Debug("LIST", "query", fmt.Sprintf("%+v", d))

	orgRes := organisation.List(ctx, cc)
	if orgRes.HasError() {
		return nil, orgRes.Error()
	}

	for _, org := range *orgRes.Payload() {
		d.StreamListItem(ctx, org)

	}
	return nil, nil
}

func getOrg(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log := plugin.Logger(ctx)
	cc := getClient()

	id := d.EqualsQuals["id"].GetStringValue()

	log.Debug("GET", "query", fmt.Sprintf("%+v", d))

	res := organisation.Get(ctx, cc, id)
	if res.HasError() {
		return nil, res.Error()
	}

	return res.Payload(), nil
}
