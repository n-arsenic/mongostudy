.PHONY: create
create:
	curl -i \
	--header "Content-Type: application/json" \
	--request POST \
	--data '{ \
	"name":"test_$(shell bash -c 'echo $$RANDOM')", \
	"int_val":100, \
	"float_val": 3.40282346638528859811704183484516925440e+38, \
	"any_list": ["a","b","c"], \
	"time": "2021-12-03T18:38:00Z"}' \
	http://localhost:8080/sample

get-list:
	curl "http://localhost:8080/sample"

update:
	curl -i \
		--header "Content-Type: application/json" \
		--request PUT \
		--data '{ \
		"id":"61b3846ffb4fa691f3439528", \
		"name":"UPDATED NAME 1", \
		"any_list_items": {"10":"B"}, \
		"category": "qwqwqwqwqwqwqwqw"}' \
		http://localhost:8080/sample
