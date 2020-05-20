echo "[1] Generating entities from ent schema"
entc generate ./ent/schema
echo "[2] Syncing entities with graphQL API"
go build -o syncbin automation/*
./syncbin >/dev/null
echo "[3] Fixing artifacts"
# shellcheck disable=SC2164
cd api/graph
sed -i '.old' 's/"errors"/"github.com\/minskylab\/collecta\/errors"/g' *.go
echo "[4] Formating code"
gofmt -w *.go
goimports -w *.go
echo "[5] Cleaning"
rm *.old