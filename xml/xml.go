package main

import (
	"fmt"
	"encoding/xml"
)


type  plist struct {
	Version string `xml:"-version"`
	Dict    struct {
		Key   string `xml:"key"`
		Array struct {
			Dict struct {
				Key   []string `xml:"key"`
				Array struct {
					Dict struct {
						Key    []string `xml:"key"`
						String []string `xml:"string"`
					} `xml:"dict"`
				} `xml:"array"`
				Dict struct {
					Key    []string `xml:"key"`
					String []string `xml:"string"`
				} `xml:"dict"`
			} `xml:"dict"`
		} `xml:"array"`
	} `xml:"dict"`
}




func main(){
	xmlstr :=  &plist{}
	xmlstr.Version = "1.0"

	xmlstr.Dict.Key = "items"

	xmlstr.Dict.Array.Dict.Key = []string{  "assets", "metadata"}

	xmlstr.Dict.Array.Dict.Array.Dict.Key = []string{ "kind", "url"}
	xmlstr.Dict.Array.Dict.Array.Dict.String = []string{ "software-package", "https://app.lggame.co/LeboVip.ipa"}

	xmlstr.Dict.Array.Dict.Dict.Key = []string{  "bundle-identifier", "bundle-version", "kind", "title"}
	xmlstr.Dict.Array.Dict.Dict.String = []string{ "com.leboapp.vipgame", "1.0.0", "software", "lebo"}



	xml_header :=`<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">`

	json_str,err := xml.Marshal(xmlstr)
	fmt.Println( xml_header,string(json_str),err)
}
