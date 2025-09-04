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

getmodels:
	apt install -y git git-lfs
	git lfs install

	mkdir -p models
	for i in $(seq 1 10); do
		git clone https://huggingface.co/qywok/stock_models_$i
		cd stock_models_$i && git lfs pull && cd ..
		mv stock_models_$i/*.onnx models/
		rm -rf stock_models_$i
	done