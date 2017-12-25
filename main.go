package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	buy := flag.Float64("buy", 0, "buy price")
	pct := flag.Float64("pct", 0, "pct change. e.g. 0.9 is 10% down. 1.25 is 25% increase")
	sell := flag.Float64("sell", 0, "sell price. If supplied will % profit where 0.05 is 5% profit")

	flag.Parse()

	// we need the "buy" value, and either "pct" or "sell"
	if *buy == 0.0 || (*pct == 0.0 && *sell == 0.0) {
		flag.Usage()
		os.Exit(1)
	}

	if *sell == 0 {
		// calculate % profit from buy an pct
		f := *buy * *pct
		fmt.Printf("sell=%0.2f\n", f)
		return
	}

	// calculate sell price to acheive pct profit
	f := (*sell - *buy) / *buy * 100.0
	fmt.Printf("pct=%0.2f\n", f)
}
