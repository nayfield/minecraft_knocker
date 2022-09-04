# minecraft_knocker
Tool that uses firewalld to open ports only to those who know the secret url


## How to set it up

- run firewalld
- verify you aren't using 'work' zone for anything (or change the references to another unused zone)
- compile the go and run it with the HANDLER_ROUTE env var set to your 'secret'

The included ansible tasks and example service files might help.

In this case the knock daemon runs as a user who can sudo firewall-cmd, minecraft is a different user.

You also need to punch holes for both minecraft and the knocker in your external firewall, and probably want to have a persistent DNS name.

## How do your users use it

- you give them a url of http://<my ip or hostname>:25818/secret_password
- if somebody goes to that URL, it adds their IP to the 'work' zone, and they can connect to minecraft


## Why?

- Avoid exposing minecraft to the whole Internet
- Don't require users to run VPN or port knock a sequence of ports

