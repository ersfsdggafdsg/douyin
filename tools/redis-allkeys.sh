echo "keys $1" | redis-cli | sed 's/^/get /' | while read cmd; do echo -n "$cmd -> "; echo $cmd | redis-cli; done
