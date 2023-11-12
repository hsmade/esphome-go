# ESPHome-go
This repo contains a library to act as an ESPHome server.
The server part, in the case of ESPHome, is the IOT device that talks to Home-Assistant.

It can be used to create a simple way to send sensor data to Home-Assistant, 
without the need for extra (Python) code in HA.

For an example implementation of a server, using this library, see the [example-server](cmd/example-server/example-server.go).
This library is not finished, lots of stuff is missing. See the [TODO](TODO.md) list.
