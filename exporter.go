package exporter

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// SocialMedia is  an interface of feeds
type SocialMedia interface {
	Feed() []string
}

// TextFile converts SocialMedia feeds to txt file
func TextFile(u SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}

// JSONFile converts SocialMedia feeds to json file
func JSONFile(u SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()

	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	i := 1
	md := u.Feed()
	cd := make(map[int][]string)
	for _, value := range md {
		cd[i] = append(cd[i], value)
		i++
	}

	b, err := json.MarshalIndent(cd, "\n", "")

	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}

	bytesWritten, err := f.Write(b)

	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}

	fmt.Printf("wrote %d bytes\n", bytesWritten)

	return nil
}

// XMLFile converts SocialMedia feeds to XML file
func XMLFile(u SocialMedia, filename string) error {
	x, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer x.Close()

	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	md := u.Feed()
	b, err := xml.MarshalIndent(md, "\n", "\n")
	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}

	bytesWritten, err := x.Write(b)

	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}
	fmt.Printf("wrote %d bytes\n", bytesWritten)

	return nil
}

// YAMLFile converts SocialMedia feeds to YAML file
func YAMLFile(u SocialMedia, filename string) error {
	q, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer q.Close()

	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	md := u.Feed()
	b, err := yaml.Marshal(md)
	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}

	bytesWritten, err := q.Write(b)

	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}
	fmt.Printf("wrote %d bytes\n", bytesWritten)

	return nil
}
