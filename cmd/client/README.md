## Getting the Hasura Schema
Hasura allows you to [export](https://hasura.io/docs/latest/guides/export-graphql-schema/)
the GraphQL schema.

For example, using [`graphqurl`](https://github.com/hasura/graphqurl):
```shell
gq https://my-graphql-engine.com/v1/graphql \
  -H "X-Hasura-Admin-Secret: adminsecretkey" \
  --introspect > schema.graphql
```

(The admin secret key [might be needed](https://github.com/hasura/graphql-engine/issues/5450)).

I've created a Makefile target to do this.

## Writing queries
Next, you'll need to write the queries your client will use.

Put them in `genqlient.graphql`.

## Run [`genqlient`](https://github.com/Khan/genqlient)
Run
```shell
go run github.com/Khan/genqlient --init
```
