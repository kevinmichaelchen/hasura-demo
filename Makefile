.PHONY: console
console:
	hasura --project schema console

.PHONY: pull-schema
pull-schema:
	(cd cmd/client && pnpm dlx graphqurl http://localhost:8080/v1/graphql --introspect > schema.graphql)

.PHONY: gen-client
gen-client:
	go run github.com/Khan/genqlient cmd/client/genqlient.yaml