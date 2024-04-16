all: clean-dev install dev

install:
	./scripts/cities_csv.sh
	./scripts/tags_csv.sh

# DEV
dev:
	ENV=DEV ./scripts/to_db.sh
clean-dev:
	rm -rf db.sqlite datasets/tags.csv
build-dev:
	ENV=DEV go build -tags "fts5" .
re-dev: clean-dev install dev build-dev

# TEST
test:
	ENV=TEST ./scripts/to_db.sh
clean-test:
	rm -rf test.sqlite datasets/tags.csv
build-test:
	ENV=TEST go build -tags "fts5" .
re-test: clean-test install test build-test

# PROD
prod:
	ENV=PROD ./scripts/to_db.sh
clean-prod:
	./scripts/confirm.sh && rm -rf prod.sqlite datasets/tags.csv || echo "aborted"
build-prod:
	ENV=PROD go build -tags "fts5" .
re-prod: clean-prod install prod build-prod

