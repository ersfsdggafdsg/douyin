hz new
ls ../../idl/http | grep -v 'api.proto' | while read file; do
	hz update -idl ../../idl/http/$file
done
