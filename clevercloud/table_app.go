package clevercloud

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"go.clever-cloud.dev/sdk/models"
	"go.clever-cloud.dev/sdk/v2/organisation"
)

func AppTable(ctx context.Context) *plugin.Table {
	log := plugin.Logger(ctx)
	log.Trace("AppTable()")

	appTable := &plugin.Table{
		Name:        "cc_app",
		Description: "CleverCloud applications",
		Columns:     extractColumns(ctx, models.ApplicationView{}),
		List:        &plugin.ListConfig{Hydrate: listApps},
		Get:         &plugin.GetConfig{Hydrate: getApp, KeyColumns: plugin.SingleColumn("id")},
	}

	return appTable
}

func listApps(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log := plugin.Logger(ctx)
	cc := getClient()

	log.Debug("LIST", "Q", d.Quals.String())

	orgIDs, err := GetOrgsList(ctx, d)
	if err != nil {
		return nil, err
	}

	for _, orgID := range orgIDs {
		log.Debug("LIST", "query", fmt.Sprintf("%+v", orgID))
		res := organisation.ListApplications(ctx, cc, orgID)
		if res.HasError() {
			return nil, res.Error()
		}

		for _, app := range *res.Payload() {
			d.StreamListItem(ctx, app)
		}
	}
	return nil, nil
}

func getApp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	log := plugin.Logger(ctx)
	cc := getClient()

	id := d.EqualsQuals["id"].GetStringValue()

	orgIDs, err := GetOrgsList(ctx, d)
	if err != nil {
		return nil, err
	}

	for _, orgID := range orgIDs {
		log.Debug("GET", "query", fmt.Sprintf("%+v", d))

		res := organisation.GetApplication(
			ctx,
			cc, orgID,
			id,
		)
		if res.IsNotFoundError() {
			continue
		}
		if res.HasError() {
			return nil, res.Error()
		}

		return res.Payload(), nil
	}

	return nil, nil
}
