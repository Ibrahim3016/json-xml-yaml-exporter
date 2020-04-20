package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"week3/facebook"
	"week3/linkedin"
	"week3/twitter"

	"gopkg.in/yaml.v3"
)

// SocialMedia is  an interface of feeds
type SocialMedia interface {
	Feed() []string
}

func main() {

	fb := new(facebook.Facebook).Feed()
	twtr := new(twitter.Twitter).Feed()
	lnkdn := new(linkedin.Linkedin).Feed()

	fbjsonfeed, _ := json.MarshalIndent(fb, "", "")
	_ = ioutil.WriteFile("fb.json", fbjsonfeed, 0644)

	fbxmlfeed, _ := xml.MarshalIndent(fb, "", "")
	_ = ioutil.WriteFile("fb.xml", fbxmlfeed, 0644)

	fbyamlfeed, _ := yaml.Marshal(&fb)
	_ = ioutil.WriteFile("fb.yaml", fbyamlfeed, 0644)

	twtrjsonfeed, _ := json.MarshalIndent(twtr, "", "")
	_ = ioutil.WriteFile("twtr.json", twtrjsonfeed, 0644)

	twtrxmlfeed, _ := xml.MarshalIndent(twtr, "", "")
	_ = ioutil.WriteFile("twtr.xml", twtrxmlfeed, 0644)

	twtryamlfeed, _ := yaml.Marshal(&twtr)
	_ = ioutil.WriteFile("twtr.yaml", twtryamlfeed, 0644)

	lnkndjsonfeed, _ := json.MarshalIndent(lnkdn, "", "")
	_ = ioutil.WriteFile("lnkdn.json", lnkndjsonfeed, 0644)

	lnkndxmlfeed, _ := xml.MarshalIndent(lnkdn, "", "")
	_ = ioutil.WriteFile("lnkdn.xml", lnkndxmlfeed, 0644)

	lnkdnyamlfeed, _ := yaml.Marshal(&lnkdn)
	_ = ioutil.WriteFile("lnkdn.yaml", lnkdnyamlfeed, 0644)

}
