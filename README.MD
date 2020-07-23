# QuickShip 🚢
quickly deploy projects to multiple servers

## How it works
Auto creates a webhook to your project and auto rebuilds and restart whenever a push is made


### Quickstart
First create a `.env`
```sh
mkdir qs
cd qs
touch .env
vim .env # edit .env
```


start a master server
```sh
quickship -s -p 7000 // master node server
```
Do the same on slave node
```sh
quickship -s -p 8000 // slave node server
```
Now create a `proj.json` json file that looks something like this :
```jsonc
{
    "nodes": ["http://123.456.789:8000"], // slave nodes
    "gitrepo": "https://github.com/abc/def", // repo of the project you want to deploy
    "branch": "master",
    "build":"docker build -t def .", // build command (optional)
    "run": "docker-compose up --build -d", // command for starting
    "clean": "docker-compose down", // cleaning up, (runs before each build except first time,and also right before removing)
    "health": {
        "path": "/ping",
        "port": 3456,
        "interval": 60000,
        "checks": 3,
        "run": "docker-compose down"
    },
    "addfiles": [
        { "name": ".env", "url": "https://abc.example.com/.env" } // files to be added to the project (added on first run before build)
    ]
}
```
Add this deployment
```sh
quickship -p 7000 -up ./proj.json -n proj0
```
Project should now be deployed

To Check list of deployed projects:
```sh
quickship -p 7000 -l
```

When you want to stop and remove a deployment record :
```sh
quickship -p 7000 -down abc123 # supply ID to -down
```


### Server folder structure
A `.env` file must be found in the root level of the directory where quickship is ran
It's recommended to run the slave quickship server inside an empty folder since it will populate project folders inside the directory the command is ran
```sh
mkdir ~/qs
cd ~/qs
touch .env
vim .env
quickship -s -p 7000
```


### Build
```sh
go build -o bin/quickship
```

### Server
Start a master / slave server node server on port 7000
```
quickship -s -p 7000
```

### Specify Master server
The `-ms` flag specifies the master server to connect to
```
quickship -ms http://123.456.1.0 -p 7000 -l
```
### Get info of a deployment record
```sh
quickship -id abc123 -i
```


### Flags
| Flag | value | Description |
| -----|-----|-----|
| -s | NULL | add this flag to start in server mode|
| -p | int | port number of the server |
| -up | string | path of deployment json file to be added |
| -n | string | name of deployment (used together with -up)|
| -down | string | ID of the deployment record to be removed |
| -l | NULL | add this flag to list deployment records|
| -ms | string | master server url |
| -id | string | ID of a specific record |
| -i | NULL | detailed info of a specific record|


### Deployment file
| Property | Value | Description| example |
| ----- | ----- | ----- | ----- |
| nodes | []string | array of slave node address | ["http://123.456.789:7000"] |
| gitrepo | string | repo url | "https://github.com/abc/def" |
| branch | string | branch name | "master" |
| run | string | run command | "docker-compose up --build -d"|
| clean |string | clean command | "docker-compose down" |
| addfiles | []file | additional files to be added on first build | [{"name":".env","url":"https://abc.example.com/.env"}] |


### Roadmap
- [x] basic build,start and stop
- [x] list
- [x] specify master server
- [ ] Suppport platforms other than Github
- [x] health check
- [ ] better key value store
- [ ] Networking for linking components
- [ ] Built-in Loadbalancer
- [ ] Cluster mode
- [x] Add/Remove nodes dynamically
- [ ] Monitor health of node
- [ ] Improve List
- [ ] Update deployed record
- [ ] Load health check upon startup
- [ ] Allow to use name as identifier (must specify ID if contains multiple of same name)
- [ ] Self-update
- [ ] Optional health check
- [ ] More ways to check health
- [ ] Full rebuild/re-deploy all