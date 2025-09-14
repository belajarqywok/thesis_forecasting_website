run:
	npm run min:css
	npm run min:download
	npm run min:js:emiten
	npm run min:js:infographic:globals
	npm run min:js:infographic:stock_prediction
	npm run min:js:infographic:stock_technical:chart
	npm run min:js:infographic:stock_technical:table
	npm run min:js:infographic:stock_historical:chart
	npm run min:js:infographic:stock_historical:table
	go run main.go

req:
	curl -X POST "http://127.0.0.1:7860/prediction" \
     -H "Content-Type: application/json" \
     -d '{"issuer":"BBCA","days":12}'