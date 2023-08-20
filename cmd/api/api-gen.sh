hz new
for file in ../../idl/http/*thrift; do
	echo $file
	hz update -idl $file
done
