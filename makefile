run:=go build -o .out && rm .out

all: 
	cd graphic/ && $(run)
	cd grub2ext/ && $(run)
	cd logger/ && $(run)
	cd mousearea/ && $(run)
	cd pinyin-search/ && $(run)
	cd set-date-time/ && $(run)
	cd sound/ && $(run)
	cd utils/ && $(run)


update:
	sudo apt-get update && sudo apt-get install dde-go-dbus-factory go-dlib
