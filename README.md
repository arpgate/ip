# ip

Minimal service to return the client remote IP or external IP.

The service is written in GO and runs on a linux server.

Start with:  nohup go run server.go > server.log &

Also included are configuration examples for haproxy and nginx,
which we use to redirect the trafic to port 3001

the haproxy file should be places as /etc/haproxy.cfg
the nginx file should be placed in etc/nginx/sites-enabled/


You can test by pointing your browser to: http://ip.arpgate.com

The server returns your external IP, the IP used by your ISP, your IP services provider
to connect your network. This is the IP outside on your cable modem or your ADSL modem.

If you use arpgate remotely this is the IP to connect to yiur home network.



Mat, Palo Alto 8/4/2015




