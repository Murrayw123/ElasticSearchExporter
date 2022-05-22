### What is this?
A custom prometheus exporter and logger that uses the ElasticSearch API to log a bit more information about
what ElasticSearch queries are taking the most time. 

### Requires
- `Docker`
- ElasticSearch v8.2

### How do I use it?
`./run.sh`

### Configuration (Environment Variables)
- `ES_URL`: Default `localhost:9200/`
- `ES_POLL_FREQ`: Default `8` seconds

### Manually testing against a mock ES server

1. Install JSON Server `npm install -g json-server`
2. Start JSON Server `json-server -p 9200 db.json --routes routes.json`
3. Run the program `./run.sh`
