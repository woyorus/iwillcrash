env GOOS=linux GOARCH=amd64 go build
docker build --no-cache --tag woyorus/iwillcrash:latest .
rm iwillcrash
echo ""
echo "Done! Push now using 'docker push woyorus/iwillcrash:latest'"
