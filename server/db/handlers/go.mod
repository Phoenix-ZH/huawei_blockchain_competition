module handlers

go 1.13

replace db/models => ../models

require (
	db/models v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.5.4
)
