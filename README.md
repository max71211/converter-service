Simple YAML to OpenMetrics format converter
===============

Description
--------
Service Allow to convert YAML file to OpenMetric format

Run
------
- `make run` - run service locally

Build
------
- `make build` - build binary file for current OS
- `make build-all` - build many binary files

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
