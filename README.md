# hasura-demo
Demo of Dockerized Hasura.

## `schema` directory
The `schema` directory will contain all migrations and metadata.

Migrations are SQL files which show how our database schema evolves.

[Metadata](https://hasura.io/docs/latest/migrations-metadata-seeds/manage-metadata/)
are YAML files which define everything about the Hasura Server itself:
> All changes made to the Hasura instance such as tracking tables / views / 
> custom functions, creating relationships, configuring permissions, creating 
> event triggers and remote schemas, etc. are tracked by Hasura using the 
> Metadata Catalogue.
>
> The Metadata can be version controlled to keep the Server configuration 
> in-sync with your codebase.

We can have Hasura create a `schema` directory with:
```shell
hasura init schema
```

The Hasura container has a volume configuration in the Docker Compose file:
```yaml
  hasura:
    volumes:
      - ./schema:/var/schema
    environment:
      HASURA_GRAPHQL_MIGRATIONS_DIR: "/var/schema/migrations"
      HASURA_GRAPHQL_METADATA_DIR: "/var/schema/metadata"
```

Whenever you use the Hasura CLI (the `hasura` command), you must be in the
`schema` directory. Or you can use the `hasura --project` flag, e.g.,
```shell
hasura --project schema console
```