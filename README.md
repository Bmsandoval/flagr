MVP:
* RELEASE FLAGS
  * what if a flag doesn't exist yet?
  * should I consider different environments for a given flag id? dev/test/prod/etc?
* list all flags in environment
* get status of one or more flags
* get status of all flags for environment
* toggle one or more flags in an environment
* administrative functions
    * insert one or more flags
    * retrieve details of flag
    * update details of flag

Nice to have:
    can save collection of flags in a small object, investigate dynamic protobuf generation

'Feature' flag - 
    on/off
    access levels
'Global' flag -
    on/off
'Release' flag -
    controls a feature's release to dev/test/prod environments
'Experiment' flag -





Thoughts
perhaps use dynamic protobuf building to design a more dynamic flagging service


Commands:
* install protoc on your machine https://grpc.io/docs/protoc-installation/
* go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
* protoc --go-grpc_out=. --go_out=. protos/*.proto
* export $(grep -v '^#' local.env | xargs -0) && go run .