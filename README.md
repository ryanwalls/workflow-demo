# Demo of bug
See https://github.com/sclasen/swfsm/issues/113



# To reproduce bug
Disclaimer to start:  I think this is related to threading, goroutines, channels, or some other race condition type problem because the procedure below will always demonstrate a "hang" of the application, but sometimes it hangs at different points.
* Install Docker
* Clone repo into this directory: `$GOPATH/src/github.com/3dsim/workflow`
* cd into `$GOPATH/src/github.com/3dsim/workflow`
* Build the docker image: `docker build -t "workflow" .`
* Run the docker image:  `docker run --rm -e "AWS_ACCESS_KEY_ID=<Your Access Key>" -e "AWS_SECRET_ACCESS_KEY=<Your secret>" workflow fsm`

NOTE: Running this will register a workflow in a "dev" domain in SWF.

# To demonstrate working polling
Follow procedure above, but final step run:
* Run the docker image:  `docker run --rm -e "AWS_ACCESS_KEY_ID=<Your Access Key>" -e "AWS_SECRET_ACCESS_KEY=<Your secret>" workflow decider`
