package main
 
import (
	"fmt"
	"net"
	"github.com/fatih/color"
)
// Main function 
func main(){
	var name string 
	banner := `
	_____         _                    _     _   _        
	|  __ \       | |                  | |   | | | |       
	| |  \/  ___  | |      ___    ___  | | __| | | | _ __  
	| | __  / _ \ | |     / _ \  / _ \ | |/ /| | | || '_ \ 
	| |_\ \| (_) || |____| (_) || (_) ||   < | |_| || |_) |
	 \____/ \___/ \_____/ \___/  \___/ |_|\_\ \___/ | .__/ 
							| |					   
													
	
	`
	color.Green("%s",banner)
	color.Blue("GoLookup --> DNS Probing Tool ")
	color.Blue("What it looks for : ")
    color.Blue(" * CNAMES : ")
	color.Blue(" * TXT Records ")
	color.Blue(" * A and AAAA Records ")
	color.Blue(" * NS (nameserver) Records  ")
	color.Blue(" * MX Records ")
	color.Blue("\n\n Enter domain name : ")
	fmt.Scanf("%s",&name)
	cname(name)
	txtrec(name)
	aaa(name)
	nameserv(name)
	mxrec(name)
}
//Function for finding CNAME
func cname(name string ) {
	color.Red("\nCanonical Name (CNAME) ")
	color.Red("+-----------------------------------------+")
	cname, _ := net.LookupCNAME(name)
	fmt.Println("[+]",cname)
}
//function for finding txt records 
func txtrec(name string) {
	color.Red("\nTXT records ")
	color.Red("+-----------------------------------------+")
	txtrecords, _ := net.LookupTXT(name)
 
	for _, txt := range txtrecords {
		fmt.Println("[+]",txt)
	}
}
//function for finding txt A and AAAA Records
func aaa(name string) {
	color.Red("\nA and AAAA  ")
	color.Red("+-----------------------------------------+")
	iprecords, _ := net.LookupIP(name)
	for _, ip := range iprecords {
		fmt.Println("[+]",ip)
	}
}
//function for finding nameserver(s)
func nameserv(name string ) {
	color.Red("\nName Server(s) (NS)  ")
	color.Red("+-----------------------------------------+")
	nameserver, _ := net.LookupNS(name)
	for _, ns := range nameserver {
		fmt.Println("[+]",ns)
	}
}
//function for finding MX record
func mxrec(name string) {
	color.Red("\nMX ")
	color.Red("+-----------------------------------------+")
	mxrecords, _ := net.LookupMX(name)
	for _, mx := range mxrecords {
		fmt.Println("[+]",mx.Host, mx.Pref)
	}
	End:
	goto End //This Label is here intentionally 
}
