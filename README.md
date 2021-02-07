
# go.base  
  
This Repository is a base Golang HTTP(s) Server with a bunch of default boilerplate (Fx, PostgreSQL, Mocks, MakeFile, etc). 

This can be used as a scaffold for future HTTP(s) Golang Servers.

You can read a more detailed explanation of this Scaffold on [Medium](https://javin-ambridge.medium.com/scaffolding-a-new-golang-http-service-f88ab8466104), including a step by step instruction set.

## Configurations
This server is built on the [Uber-Go Fx Dependency Management system](https://github.com/uber-go/fx).

This server is setup similar to how internal Uber services are setup, as well as a few minor changes.

## Directory Structure
|Directory|Usage|
|--|--|
| .gen | This Directory is used to contain any generated code (ie. Mocks). |
| app | This Directory is used to define any Fx Options. |
| canned-requests | This directory contains a simple python script ( 'request.py' ) that can be modified slightly to hit HTTP endpoints during testing.<br><br>The corrolary of this for YARPC services is Yabs. |
| config | This Directory is used to define static YAML configurations. |
| constants | This Directory is used to define any constants (metric fields, error messages, etc). |
| controller | This Directory is used to build actual code logic.<br><br>Try to break your logic flows into controllers that you can call from the handler(s). |
| db | This Directory is used to for any logic related to your database.<br><br>Initially provided is a PostgreSQL DB setup. |
| entity | This Directory is used to any structures used in your service. |
| gateway | This Directory is used to define any functions to interact with external services, applications, etc. |
| handler | This Directory is used to define endpoint handlers (ie. Request comes in one type, and needs to leave the same type). |
| observability | This Directory is used to define any Observability Options/Setup. This would be things like Loggers, M3 Metric Scopes, etc |
| scaffold | This Directory contains the script to convert this go.base to your actual application.<br><br>This directory will be deleted after you convert the go.base application to your own. |
| secrets | This Directory is used to store your application secrets.<br><br>The secrets.go file is looking for a YAML file called 'secrets.yaml' in this directory. |
| server | This Directory is used to setup your HTTP server. |
| utils | This Directory is used to define any Utilities. |

## How to Clone use this as a Scaffold

 1. Create a new directory named whatever you want at $GOPATH/src/github.com/GITHUB_NAME/
 2. Go into that new directory
 3. Clone this repository
 4. Go into this new repository
 5. Run the scaffold/convert.py script (python scaffold/convert.py)
 6. Fill in the input
 7. Add, Commit, and Push to your new repository.
 
 ## Questions
 
 Feel free to put issues if you would like any more boilerplate added (or add a PR), can easily be achievable. 
