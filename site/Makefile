all: restart

restart:
	docker stop devsyringe || true
	docker rm devsyringe || true
	docker build -t devsyringe-landing .
	docker run -d \
		--name devsyringe \
		--network personal-site_default \
		-p 5143:5143 \
		--restart=always \
		devsyringe-landing
