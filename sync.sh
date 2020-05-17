echo "Generating entities from ent schema"
entc generate ./ent/schema
echo "Syncing entities with graphQL API"
go build -o syncbin automation/*
./syncbin