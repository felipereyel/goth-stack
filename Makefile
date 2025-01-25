all: statics templ

statics:
	curl https://cdn.tailwindcss.com/3.4.16 --output src/embeded/statics/tailwind.js
	curl https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js --output src/embeded/statics/htmx.js

templ:
	templ generate
