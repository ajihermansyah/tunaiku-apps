# tunaiku-apps

## Overview

### Application Requirement
* [Go Lang 1.7.1](https://golang.org/dl/) as Tools/Programming language
* [MongoDB v2.6.10](https://docs.mongodb.com/)
* [Bee version v1.9.1]
* [Beego Framework 1.7.1]

### Installation
#### Set Environment Variables (GOPATH)
The GOPATH environment variable specifies the location of your workspace. It is likely the only environment variable you'll need to set when developing Go code.

To get started, create a workspace directory and set GOPATH accordingly. Your workspace can be located wherever you like, but we'll use $HOME/go in this project.
$ export GOPATH=$HOME/go


#### Clone airport-shuttle-bus-api project from repository (this repository)
The project must be cloned under $GOPATH/src directory.

$ cd $GOPATH/src
$ git clone [repository URL.git]


### Application Configuration
In conf/app.conf, you must provide port to be used and database configuration
#### Application Port
port = 8080


#### Database Configuration
appname = tunaiku
httpport = 8080
runmode = dev
copyrequestbody = true
sessionon = true

mgo_host=127.0.0.1:27017
mgo_database=tunaiku
mgo_username=
mgo_password=
mgo_host_seeder="localhost:27017"


#### Run APPS
bee run 
access localhost:8080/stock on browser