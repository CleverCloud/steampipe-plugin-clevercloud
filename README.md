![image](https://hub.steampipe.io/images/plugins/turbot/clevercloud-social-graphic.png)

# CleverCloud Plugin for Steampipe

Use SQL to query platform services, apps and addons from your Clever Cloud account.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/clevercloud)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/clevercloud/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/CleverCloud/steampipe-plugin-clevercloud/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install clevercloud
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/clevercloud#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/clevercloud#configuration).

Run a query:

```sql
select
  id,
  name
from
  cc_apps;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/CleverCloud/steampipe-plugin-clevercloud.git
cd steampipe-plugin-clevercloud
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/clevercloud.spc
```

Try it!

```shell
steampipe query
> .inspect clevercloud
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/CleverCloud/steampipe-plugin-clevercloud/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [PagerDuty Plugin](https://github.com/CleverCloud/steampipe-plugin-clevercloud/labels/help%20wanted)