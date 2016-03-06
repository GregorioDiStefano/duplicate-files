dd if=/dev/urandom of=test1/big bs=1024 count=10240
cp test1/big test1/big2

size=$(expr 1024 \* 1024)

rm test1/testfile.1

for ((i=1; i<=$size; i++));
do
   printf '\xaa\xbb\xcc\xdd' >> test1/testfile.1
done
