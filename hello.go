package main

import (
	"fmt"
	"log"

	"gopkg.in/ldap.v2"
)

func main() {
	auth()
}

func auth() {
	// The username and password we want to check
	username := "someuser"
	password := "userpassword"

	bindusername := "readonly"
	bindpassword := "password"

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 389))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	l.Start()

	// First bind with a read only user
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
	}

	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	// Bind as the user to verify their password
	err = l.Bind(userdn, password)
	if err != nil {
		log.Fatal(err)
	}

	// Rebind as the read only user for any further queries
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
	}
}
