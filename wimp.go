// wimp: What Is My iP address
package main

import (
        "fmt"
        "io/ioutil"
        "net"
        "net/http"
        "os"
        "strings"
)

const url = "http://ifconfig.me/ip"

func main() {
        resp, err := http.Get(url)
        if err != nil {
                fmt.Printf("[!] GET failed: %s\n", err.Error())
                os.Exit(1)
        }
        defer resp.Body.Close()
        baddr, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                fmt.Printf("[!] failed to read response: %s\n", err.Error())
                os.Exit(1)
        }
        addr := strings.TrimSpace(string(baddr))
        fmt.Printf("[+] your IP address is %s", addr)

        hosts, err := net.LookupAddr(addr)
        if err != nil {
                fmt.Printf("\n[!] LookupAddr failed: %s\n", err.Error())
                os.Exit(1)
        }
        fmt.Printf(" (%s)\n", strings.Join(hosts, ", "))
}
