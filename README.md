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

The `schema` directory was created with:
```shell
hasura init schema
```
(This only needs to happen once).

### Developer Workflow

Developers will need to run Hasura locally to create any migrations.

The Docker volume is critical for allowing your local Hasura Server create
and update migrations and metadata files.
```yaml
  hasura:
    volumes:
      - ./schema:/var/schema
    environment:
      HASURA_GRAPHQL_MIGRATIONS_DIR: "/var/schema/migrations"
      HASURA_GRAPHQL_METADATA_DIR: "/var/schema/metadata"
```

## Console
To start the Hasura Console with the CLI, run:
```shell
make console
```

## GraphQL
Hasura generates a very powerful [GraphQL API](https://hasura.io/docs/latest/getting-started/how-it-works/index/#tracking-tables--schema-generation).

Here are the most basic examples:

### Insert
To create a new animal:
```graphql
mutation CreateAnimal {
  insert_animal(objects: {
    species: "Canis lupus"
  }) {
    returning {
      id
      species
    }
  }
}
```

### Query
To list all animals:
```graphql
{
  animal {
    id
    species
  }
}
```