test: opts :=
test: v :=
test:
	rm -rf ../../.allure_results || true
	python3 -m pytest $(opts) $(if $(v),-o log_cli=True)

allure: web :=
allure:
	allure $(if $(web),serve,generate) ../../.allure_results --output ../../.allure_report --host 127.0.0.1

bootstrap:
	pip3 install -U pip
	pip3 install -U -r requirements.txt
