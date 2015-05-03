package ucl

//go:generate ragel -Z ucl.rl
//go:generate sh -c "ragel -Z -S value -V -p ucl.rl | dot -Tpng > parser.png"
