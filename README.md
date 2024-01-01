# fix-a-pix-puzzle-minisat

go run main.go
cp tvars check_art
sed -e '1d' out
mv out check_art

cd check_art
go run main.go
