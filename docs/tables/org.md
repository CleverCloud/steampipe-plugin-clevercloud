# Table: cc_org

`cc_org` define a Clever Cloud organization where you can manage applications, team members and managed services.

See the [Organization documentation](https://www.clever-cloud.com/doc/account/administrate-organization/).

## Examples

### Basic info

```sql
select
  id,
  name,
  contry,
  vat,
  avatar
from
  cc_org;
```