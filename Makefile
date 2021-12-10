all:
	docker-compose up --build -d
	docker-compose exec app /bin/bash

fclean:
	docker-compose down --volumes --remove-orphans

re: fclean all

.PHONY: all in clean fclean re