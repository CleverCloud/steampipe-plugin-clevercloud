---
organization: CleverCloud
category: ["public cloud"]
icon_url: "https://www.clever-cloud.com/app/themes/cc-wp-theme/assets/img/brand-assets/logo_on_white.svg"
brand_color: "#d74d4e"
display_name: Clever Cloud
name: clevercloud
description: Steampipe plugin for querying Clever Cloud applications, addons and more.
og_description: Query Clever Cloud with SQL! Open source CLI. No DB required.
og_image: "https://www.clever-cloud.com/app/themes/cc-wp-theme/assets/img/brand-assets/logo_on_white.svg"
---

# Clever Cloud + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Clever Cloud](https://clever-cloud.com) Clever Cloud handles all the ops work while you focus on your true business value: coding.

For example:

```sql
select
  id,
  name
from
  cc_apps;
```

```
+------------------------------------------+--------------+
| id                                       | name         |
+------------------------------------------+--------------+
| app_xd252ccf-af5b-4d9f-a7d7-ca1fd4d0a2ed | querier      |
+------------------------------------------+--------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/clevercloud/tables)**

## Get started

### Install

Download and install the latest Clever Cloud plugin:
vv
```bash
steampipe plugin install clevercloud
```

### Credentials

Credentials are read from `~/.config/clever-cloud/clever-tools.json`.

This file can be filled using `clever-tools login` command.

See [clever-tools](https://github.com/CleverCloud/clever-tools) documentation.


### Configuration

Installing the latest alicloud plugin will create a config file (`~/.steampipe/config/clevercloud.spc`) with a single connection named `clevercloud`:

```hcl
connection "clevercloud" {
}
```

## Get involved

* Open source: https://github.com/CleverCloud/steampipe-plugin-clevercloud
* Community: [Slack Channel](https://steampipe.io/community/join)
