all: Setup Downloads Generate

Setup:
	go get -d github.com/go-swagger/go-swagger/cmd/swagger
	go install github.com/go-swagger/go-swagger/cmd/swagger
	go mod tidy
Downloads:
	curl -sSL https://esi.evetech.net/_latest/swagger.json -o esi.json
	curl -sSL https://esi.evetech.net/swagger.json -o meta.json
	-swagger mixin -q esi.json meta.json -o swagger.json
Generate:
	swagger generate client -f swagger.json -c esi -m models --default-scheme=https --with-enum-ci --skip-validation \
	--tags "Alliance" \
	--tags "Assets" \
	--tags "Bookmarks" \
	--tags "Calendar" \
	--tags "Character" \
	--tags "Clones" \
	--tags "Contacts" \
	--tags "Contracts" \
	--tags "Corporation" \
	--tags "Dogma" \
	--tags "Faction Warfare" \
	--tags "Fittings" \
	--tags "Fleets" \
	--tags "Incursions" \
	--tags "Industry" \
	--tags "Insurance" \
	--tags "Killmails" \
	--tags "Location" \
	--tags "Loyalty" \
	--tags "Mail" \
	--tags "Market" \
	--tags "Opportunities" \
	--tags "Planetary Interaction" \
	--tags "Routes" \
	--tags "Search" \
	--tags "Skills" \
	--tags "Sovereignty" \
	--tags "Status" \
	--tags "Universe" \
	--tags "User Interface" \
	--tags "Wallet" \
	--tags "Meta" \
	--tags "Wars"
	go get -u -d ./esi/...
	go get -u -d ./models/...
	go run scripts/generator.go scripts/sde.go
	go mod tidy

#cat .\swagger.json | jq -r '[.paths | to_entries | .[].value | to_entries | .[].value.tags[0]] | sort | unique | join(\", \")'
