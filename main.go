//go:generate pigeon -o main_gen.go host.peg

package main

import (
	"log"
	"os"
)

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

func toStringSlice(v interface{}) []string {
	if v == nil {
		return nil
	}
	vv := v.([]interface{})
	strs := make([]string, len(vv))
	for i := range vv {
		s, ok := vv[i].(string)
		if !ok {
			continue
		}
		strs = append(strs, s)
	}
	return strs
}

func toASTHostLineSlice(v interface{}) []ASTHostLine {
	if v == nil {
		return nil
	}
	vv := v.([]interface{})
	lns := make([]ASTHostLine, 0, len(vv))
	for i := range vv {
		l, ok := vv[i].(ASTHostLine)
		if !ok {
			continue
		}
		lns = append(lns, l)
	}
	return lns
}

func main() {
	log.Println("Parsing hosts")
	h, err := ParseFile("host")
	hlns := h.(ASTHosts)

	if err != nil {
		log.Println("Got Error:", err)
		os.Exit(1)
	}
	log.Printf("Hostlines: %T\n\n%v\n", hlns, hlns.StringWOComments())
}
