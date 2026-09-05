package graphql

const (
	Query    = "query"
	Mutation = "mutation"
	Subscription = "subscription"
)

const (
	IntrospectionQuery = `
		query IntrospectionQuery {
			__schema {
				queryType { name }
				mutationType { name }
				subscriptionType { name }
				types {
					kind
					name
					description
				}
			}
		}
	`
)
