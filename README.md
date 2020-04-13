
# go.base  
  
This Repository is a base Golang HTTP Server with a bunch of default setting up done. 

This can be used as a scaffold for future HTTP Golang Servers.

## Configurations
This server is built on the [Uber-Go Fx Dependency Management system](https://github.com/uber-go/fx).

This server is setup similar to how internal Uber services are setup, as well as a few minor changes.

## Directory Structure
|Directory|Usage|
|--|--|
| app | This Directory is used to define any Fx Options. |
| canned-requests | This directory contains a simple python script ( 'request.py' ) that can be modified slightly to hit HTTP endpoints during testing.<br><br>The corrolary of this for YARPC services is Yabs. |
| config | This Directory is used to define static YAML configurations. |
| constants | This Directory is used to define any constants (metric fields, error messages, etc). |
| controller | This Directory is used to build actual code logic.<br><br>Try to break your logic flows into controllers that you can call from the handler(s). |
| db | This Directory is used to for any logic related to your database. |
| entity | This Directory is used to any structures used in your service. |
| gateway | This Directory is used to define any functions to interact with external services, applications, etc. |
| handler | This Directory is used to define endpoint handlers (ie. Request comes in one type, and needs to leave the same type). |
| observability | This Directory is used to define any Observability Options/Setup. This would be things like Loggers, M3 Metric Scopes, etc |
| server | This Directory is used to setup your HTTP server. |
| utils | This Directory is used to define any Utilities. |

## How to Clone use this as a Scaffold

 1. Create a new directory named whatever you want at $GOPATH/src/github.com/GITHUB_NAME/
 2. Go into that new directory
 3. Clone this repository
 4. Rename the go.base directory to be the same name as what you created in step 1
 5. Rename all instances of 'go.base' in all of the go files, with your new service name (find . -type f -name "*.go" -print0 | xargs -0 sed -i '' -e 's/go\.base/**NAME_FROM_STEP_1**/g')
 6. Rename all instances of 'go.base' in the Makefile with your new service name (find . -type f -name "Makefile" -print0 | xargs -0 sed -i '' -e 's/go\.base/**NAME_FROM_STEP_1**/g')
 7. Update the Readme
 8. Create a Repository on GitHub
 9. Override the origin URL for this new repository (git remote set-url origin ...) 
 10. Push your first commit 
 11. Done, you now have your service scaffolded
