package internal

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent --feature sql/lock,sql/modifier,sql/upsert,privacy,entql ../schema
