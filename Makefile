image:
	docker build -t summarize-covid19-stats:latest -f Dockerfile .

container:
	docker run --rm -p 8899:8899 --env-file ./local.env --name summarize-covid19-stats summarize-covid19-stats:latest