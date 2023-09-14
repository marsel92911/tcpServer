# tcpServer

### 1. Choosing a Proof of Work Algorithm:

I chose the Hashcash algorithm.
The algorithm was chosen because of its prevalence, ease of implementation and understanding.
The algorithm also allows you to adjust the complexity of the calculation for the client.


### 2. Information on program launches:
First you need to run a contrainer from the server.<br>
To launch a docker container with a tcp server application, run the command

`make run-server`

Then you need to run a contrainer with the client.<br>
To launch a docker container with a client application, you need to run the command

`make run-client`

The TCP port of the server can be set using an environment variable, using the command:

`make run-server -e TCP_SERVER_PORT="12345"`

Changes to the complexity of PoW needed can be made using the environment variable `TCP_SERVER_TARGET_BITS` (default value is 18):

`make run-server -e TCP_SERVER_TARGET_BITS="20"`
