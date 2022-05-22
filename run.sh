docker build -t es-task-prom-log-exporter . && docker run -d -e ES_URL="{$ES_URL}" \
-e ES_POLL_FREQ="{$ES_POLL_FREQ}" \
-p 2112:2112 \
es-task-prom-log-exporter