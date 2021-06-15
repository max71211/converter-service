Simple YAML to OpenMetrics format converter
===============

Description
--------
Service Allow to convert YAML file to OpenMetric format

Run
------
make run

Build
------
make build 
make build-all

Configuration
------------
- `HOST` - host for service
- `PORT` - port for service
- `YAML_FILE_PATH` - YAML file path

API
---
### GET /
Return Index Handler string

### GET /metrics
Return OpenMetrics converted string
