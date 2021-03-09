package main

import (
	"fmt"
	"net"
	"flag"
	"bufio"
	"os"
)


func logo() {
	fmt.Println(`

       __               
  ____/ /___  __________
 / __  / __ \/ ___/ ___/
/ /_/ / / / (__  ) /__  
\__,_/_/ /_/____/\___/ 

     @channyeinwai(1.0)

`)

}

func main() {

	var domains []string

	logo()

	if flag.NArg() > 0 {
		domains = []string{flag.Arg(0)}

	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			domains = append(domains , sc.Text())
		}

		if err := sc.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
		}
	}

	var domain string

	for _, domain = range domains{
		ips, _ := net.LookupIP(domain)
		for _,ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				fmt.Println(ipv4)
			}
		}
	}
}
