go test ./... -coverprofile=cover.out -v || exit 1
go tool cover -html=cover.out -o cover.html
rm cover.out
