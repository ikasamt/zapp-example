PROJECT_ID := XXXXXXXXX
DB_INSTANCE := XXXXXXXXX
EMAIL := XXX@XXX

info:
	@echo ""
	@echo "  gcloud app deploy --project $(PROJECT_ID)"
	@echo "  gcloud app deploy queue.yaml --project $(PROJECT_ID)"
	@echo "  gcloud app deploy cron.yaml --project $(PROJECT_ID)"
	@echo "  gcloud app browse --project $(PROJECT_ID)"
	@echo "  gcloud sql connect $(DB_INSTANCE) --project $(PROJECT_ID) "
	@echo ""

goemon:
	goemon --

build: js gss zapp-jam
	@echo "build ..."

dev: build
	dev_appserver.py app/app.yaml 2>&1

clean:
	rm -rf zzz-*
	rm -rf models/zzz-*

gss:
	@echo "COMPILE css START..."
	@docker run -i -v `PWD`/assets:/opt/assets \
		ikasamt/google-closure-tools \
		bash -c 'java -jar closure-stylesheets.jar /opt/assets/gss/*.css' \
		> public/css/all.css
	@echo "COMPILE css DONE"

js:
	@echo "COMPILE js START..."
	@docker run -i -v `PWD`/assets:/opt/assets \
		ikasamt/google-closure-tools \
		bash -c 'java -jar closure-compiler.jar /opt/assets/js/*.js' \
		> public/js/all.js
	@echo "COMPILE js DONE"

zapp-jam: clean
	zapp-jam models


