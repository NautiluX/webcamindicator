all: build

build:
	go build

icons:
	make -C icon

install:
	cp webcamindicator /usr/local/bin/
	cp icon/cam-white.png /usr/share/icons/
	cat webcamindicator.desktop | sed s@DEVICE@$(DEVICE)@ > /usr/share/applications/webcamindicator.desktop

autostart:
	cat webcamindicator.desktop | sed s@DEVICE@$(DEVICE)@ > ~/.config/autostart/webcamindicator.desktop
