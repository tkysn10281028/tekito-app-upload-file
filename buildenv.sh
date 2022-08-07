while true
do
echo build frontend?[y/n]
read str
if [ $str = y ]; then
    cd ~/uploadFilePJ/front
    npm run build
    break
fi
if [ $str = n ]; then
    break
fi
done
cd ~/uploadFilePJ
go build
./uploadFilePJ