#WWWHOME=	/var/www/htdocs
WWWHOME=/var/www/vhost/oldcode.org/htdocs/

deploy:
	cp -r static ${WWWHOME}
	cp -r URT_RADIO_RAW ${WWWHOME}/static/

.PHONY= deploy

