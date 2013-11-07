// http://www.let.rug.nl/~kleiweg/go/xpath/
package main

import (
	"fmt"
	"github.com/moovweb/gokogiri/xml"
	//"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("xpath.xml")
	util.CheckErr(err)

	doc, err := xml.Parse(data, nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)
	defer doc.Free()

	n, err := doc.Root().Search(`/probe/header/@vendor`)
	fmt.Println("Vendor:")
	fmt.Println(n)
	util.CheckErr(err)
}
