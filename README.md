# Golookup
GoLookup is a simple tool written in golang , which looks for CNAME(s) ,A and AAAA Records , TXT Records , NameServer(s) / MX Record of any domain 

[+] CNAMES(s) : CNAME is a Canonical Name Record or Alias Record. A type of resource record in the Domain Name System (DNS), that specifies that one domain name is an alias of another canonical domain name.(src = wikipedia)

[+] A and AAAA Record : An A and AAAA record are actually primary DNS records.  They associate a domain name with a specific IP address.

[+] TXT Record : A TXT record (short for text record) is a type of resource record in the Domain Name System (DNS) used to provide the ability to associate arbitrary text with a host or other name, such as human readable information about a server, network, data center, or other accounting information.(src = wikipedia)

[+] MX(Mail Exchanger) Record : A mail exchanger record (MX record) specifies the mail server responsible for accepting email messages on behalf of a domain name. It is a resource record in the Domain Name System (DNS). It is possible to configure several MX records, typically pointing to an array of mail servers for load balancing and redundancy. (src = wikipedia)

[+] NameServer(s) : Nameserver is a server on the internet specialized in handling queries regarding the location of a domain name's various services. Nameservers are a fundamental part of the Domain Name System (DNS). They allow using domains instead of IP addresses. (src = wikipedia)

# How to Use (Example) :

D:\GO\src\github.com\devanshwolf\golang> go run golookup.go

        _____         _                    _     _   _
        |  __ \       | |                  | |   | | | |
        | |  \/  ___  | |      ___    ___  | | __| | | | _ __
        | | __  / _ \ | |     / _ \  / _ \ | |/ /| | | || '_ \
        | |_\ \| (_) || |____| (_) || (_) ||   < | |_| || |_) |
         \____/ \___/ \_____/ \___/  \___/ |_|\_\ \___/ | .__/
                                                                                                       


GoLookup --> DNS Probing Tool
What it looks for :
 * CNAMES :
 * TXT Records
 * A and AAAA Records
 * NS (nameserver) Records
 * MX Records


 Enter domain name :
facebook.com

Canonical Name (CNAME)

+-----------------------------------------+

[+] facebook.com.

TXT records

+-----------------------------------------+

[+] v=spf1 redirect=_spf.facebook.com

A and AAAA

+-----------------------------------------+

[+] 2a03:2880:f12f:183:face:b00c:0:25de
[+] 31.13.79.35

Name Server(s) (NS)

+-----------------------------------------+

[+] &{a.ns.facebook.com.}
[+] &{b.ns.facebook.com.}

MX

+-----------------------------------------+

[+] msgin.vvv.facebook.com. 10


# Installation

1 - Install Golang(https://golang.org/doc/install).

2 - SetUp Proper paths.

3 - Install "github.com/fatih/color" Library(via go get github.com/fatih/color). 

4 - Compile the Code .

5 - Ready to run..!

# Contact : 
https://twitter.com/devanshwolf

credits : http://www.golangprograms.com/ :)
