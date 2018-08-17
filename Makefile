
default:
	ldapsearch -x -b "dc=beispiel,dc=de" "uid=pgarrett"

run:
	docker run --rm -it -p 389:389 \
		-e ORGANISATION_NAME="Beispiel gmbh" \
		-e SUFFIX="dc=beispiel,dc=de" \
		-e ROOT_PW="geheimnis" \
		pgarrett/ldap-alpine

clean:
