# fix-a-pix-puzzle-minisat

go run main.go > tvars
cp tvars check_art
cd check_art
sed -e '1d' tvars
