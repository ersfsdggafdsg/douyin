ls ../../idl/rpc/ | while read file; do 
	echo ../../rpc/$file
	kitex -I ../../idl/rpc ../../idl/rpc/$file
done
