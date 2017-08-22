## Dockerized, lightweight, swarm-ready fail2ban log monitor for StatsD written in Go


#### [DockerHub](https://hub.docker.com/r/garugaru/fail2ban-monitor/)

#### [Example Dashboard](https://snapshot.raintank.io/dashboard/snapshot/teGUakGAVk9BDRDpivR0Zbph6IhhJfYk)


### Deployment example (compose.yml):

	 version: "3.3"
	 
	 services:
	 
	   fail2ban-monitor:
		   image: garugaru/fail2ban-monitor
		   volumes:
			 - type: bind
			   source: /var/log
			   target: /logs
		   environment:
			 - STATSD_ADDRESS=127.0.0.1:8125
			 - STATSD_PREFIX=fail2ban
		   deploy:
			 mode: global
			 restart_policy:
			   condition: on-failure
			   delay: 5s
			   max_attempts: 3
	 

