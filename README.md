### What is this?

A custom prometheus exporter and logger that uses the ElasticSearch API to log a bit more information about
what ElasticSearch queries are taking the most time.

### Requires

- `Docker`
- `ElasticSearch >= v8.2`

### How do I use it?

`./run.sh`

### What does the log message look like?

`{"level":"info","ts":1653223975.9755397,"caller":"ElasticSearchExporter/main.go:72","msg":"timestamp: 2022-05-22 20:52:55, task_id: 464, running_time_ms: 13, action: indices:data/read/search, description: indices[test], types[test], search_type[QUERY_THEN_FETCH], source[{\"query\":...}]"}
`
`

### Configuration (Environment Variables)

- `ES_URL`: Default `localhost:9200/`
- `ES_POLL_FREQ`: Default `8` seconds

### Manually testing against a mock ES server

1. Install JSON Server `npm install -g json-server`
2. Start JSON Server `json-server -p 9200 db.json --routes routes.json`
3. Run the program `./run.sh`
