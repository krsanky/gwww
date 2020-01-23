WWWHOME=/var/www/htdocs

deploy:
	cp -r static ${WWWHOME}
	cp static/favicon.ico ${WWWHOME}

.PHONY= deploy

