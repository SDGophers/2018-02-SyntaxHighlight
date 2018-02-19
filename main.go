//go:generate pigeon -o main_gen.go host.peg

package main

import (
	"log"
	"os"
	"fmt"
	"github.com/fatih/color"
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
	if err != nil {
		log.Fatal(err)
	}
	hlns,  ok := h.(ASTHosts)
	if !ok {
		log.Fatalf("No tokens received\n%v", h)
	}

	blue := color.New(color.FgBlue)
	green := color.New(color.FgGreen)

	for _, v := range hlns {
		//fmt.Printf("Line %d: %t\n", k, v)
	P:
		switch val := v.(type){
		case ASTComment:
			blue.Printf("#" + val.String())
		case ASTHost:
			green.Printf(val.String())
			if val.comment != nil && val.comment.Comment != "" {
				fmt.Printf("\t")
				v = *val.comment
				goto P
			}
		case ASTBlank:
			fmt.Printf(v.String())
		default:
			log.Fatalf("Unexpected type %t ", v)
		}
	}

	if err != nil {
		log.Println("Got Error:", err)
		os.Exit(1)
	}
	//log.Printf("Hostlines: %T\n\n%v\n", hlns, hlns.StringWOComments())
}
