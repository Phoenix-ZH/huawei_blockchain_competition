module db

go 1.13

replace server.com/config => ../config

//replace handlers => ./handlers

require (
	go.mongodb.org/mongo-driver v1.5.4
	server.com/config v0.0.0-00010101000000-000000000000
)
