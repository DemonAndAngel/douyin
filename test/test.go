package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Import struct {
	Domain string `json:"domain"`
	ExpirationDate float64 `json:"expirationDate,omitempty"`
	HostOnly bool `json:"hostOnly"`
	HTTPOnly bool `json:"httpOnly"`
	Name string `json:"name"`
	Path string `json:"path"`
	SameSite string `json:"sameSite"`
	Secure bool `json:"secure"`
	Session bool `json:"session"`
	StoreID string `json:"storeId"`
	Value string `json:"value"`
	ID int `json:"id"`
}
type Export struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Domain string `json:"domain"`
	Path string `json:"path"`
	Expires int `json:"expires"`
	Size int `json:"size"`
	HTTPOnly bool `json:"httpOnly"`
	Secure bool `json:"secure"`
	Session bool `json:"session"`
	SameSite string `json:"sameSite"`
	Priority string `json:"priority"`
	SameParty bool `json:"sameParty"`
	SourceScheme string `json:"sourceScheme"`
	SourcePort int `json:"sourcePort"`
}

func main() {
	im, err := ReadAll("./import.tmp")
	if err != nil {
		panic(err)
	}
	imports := []Import{}
	json.Unmarshal(im, &imports)
	cookies := []Export{}
	for _, i := range imports {
		cookies = append(cookies, Export{
			Name:         i.Name,
			Value:        i.Value,
			Domain:       i.Domain,
			Path:         i.Path,
			Expires: int(i.ExpirationDate),
			Size:         0,
			HTTPOnly:     i.HTTPOnly,
			Secure:       i.Secure,
			Session:      i.Session,
			SameSite:     i.SameSite,
			Priority:     "Medium",
			SameParty:    false,
			SourceScheme: "",
			SourcePort:   0,
		})
	}
	exports := map[string][]Export{
		"cookies": cookies,
	}
	// 写入
	f, err := os.Create("export.tmp")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, _ := json.Marshal(exports)
	f.Write(b)
	fmt.Println("done")
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}