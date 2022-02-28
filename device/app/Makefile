SHA := master

restore:
	git reset --hard ${SHA}
	pip install -r requirements.txt

patch:
	sudo python3 patch.py

start:
	source bin/activate
	python3 main.py