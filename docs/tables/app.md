# Table: cc_app

`cc_app` define a Clever Cloud application, the entity which run your code.
See the [Application documentation](https://www.clever-cloud.com/doc/administrate/apps-management/).

## Examples

### Basic info

```sql
select
  id, 
  name,
  description,
  state,
  zone
from
  cc_app;
```