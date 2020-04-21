package main

import (
	exporter "week3"
	"week3/facebook"
	"week3/linkedin"
	"week3/twitter"
)

func main() {
	txt := (exporter.TextFile)
	son := (exporter.JSONFile)
	exp := (exporter.XMLFile)

	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdn := new(linkedin.Linkedin)

	err := txt(twtr, "twtrdata.txt")
	err = txt(fb, "fbdata.txt")
	err = txt(lnkdn, "linkdn.txt")

	if err != nil {
		panic(err)
	}

	m := son(twtr, "twtrdata.json")
	m = son(fb, "fbdata.json")
	m = son(lnkdn, "linkdn.json")

	if m != nil {
		panic(err)
	}

	v := exp(twtr, "twtrdata.xml")
	v = exp(fb, "fbdata.xml")
	v = exp(lnkdn, "linkdn.xml")

	if v != nil {
		panic(err)
	}

}
