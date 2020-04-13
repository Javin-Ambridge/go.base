
# go.base  
  
This Repository is a base Golang HTTP Server with a bunch of default setting up done. 

This can be used as a scaffold for future HTTP Golang Servers.

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
| db | This Directory is used to for any logic related to your database. |
| entity | This Directory is used to any structures used in your service. |
| handler | This Directory is used to define endpoint handlers (ie. Request comes in one type, and needs to leave the same type). |
| gateway | This Directory is used to define any functions to interact with external services, applications, etc. |
| observability | This Directory is used to define any Observability Options/Setup. This would be things like Loggers, M3 Metric Scopes, etc |
| scaffold | This Directory contains the script to convert this go.base to your actual application. |
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
