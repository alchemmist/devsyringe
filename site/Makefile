all: restart

restart: 
	docker stop devsyringe
	docker rm devsyringe
	docker build -t devsyringe-landing .
	docker run -d --name devsyringe --network personal-site_default -p 5143:5143 devsyringe-landing
