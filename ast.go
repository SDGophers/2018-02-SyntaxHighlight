package main

import (
	"net"
	"strings"
)

type ASTHostLine interface {
	Line() int
	String() string
}

type ASTComment struct {
	Comment string
	line    int
}

func (c ASTComment) String() string { return "#" + c.Comment + "\n" }
func (c ASTComment) Line() int      { return c.line }

type ASTBlank struct {
	Count int
	line  int
}

func (c ASTBlank) String() string { return strings.Repeat("\n", c.Count) }
func (c ASTBlank) Line() int      { return c.line }

type ASTHost struct {
	IP      net.IP
	Aliases []string
	line    int
}

func (h ASTHost) String() string {
	return h.IP.String() + "\t" + strings.Join(h.Aliases, " ") + "\n"
}
func (h ASTHost) Line() int { return h.line }

type ASTHosts []ASTHostLine

func (h ASTHosts) String() string {
	var strs = make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, "")
}

func (h ASTHosts) StringWOComments() string {
	var strs = make([]string, len(h))
	for i := range h {
		if _, ok := h[i].(ASTComment); ok {
			strs[i] = "\n"
			continue
		}
		strs[i] = h[i].String()
	}
	return strings.Join(strs, "")
}
