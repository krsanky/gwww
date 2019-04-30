deploy:
	cp -r static /var/www/htdocs/
	cp -r URT_RADIO_RAW /var/www/htdocs/static/

.PHONY= deploy

