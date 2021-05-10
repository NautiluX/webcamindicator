# Webcam Indicator Application for Linux

Indicates with a tray icon if the specified webcam is in use or not. Excludes obs processes. Based on https://github.com/getlantern/systray .

## Usage

```
webcamindicator /dev/video<x>
```

Specify the device that reflects your webcam.

## Build

```
make
```

## Install

```
sudo make DEVICE=/dev/video<x> install
```

## Autostart

```
sudo make DEVICE=/dev/video<x> autostart
```
