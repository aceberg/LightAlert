[![Main-Docker](https://github.com/aceberg/LightAlert/actions/workflows/main-docker.yml/badge.svg)](https://github.com/aceberg/LightAlert/actions/workflows/main-docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/LightAlert)](https://goreportcard.com/report/github.com/aceberg/LightAlert)
[![Maintainability](https://api.codeclimate.com/v1/badges/d5afcbba61b811fa3ddc/maintainability)](https://codeclimate.com/github/aceberg/LightAlert/maintainability)

<h1><a href="https://github.com/aceberg/LightAlert">
    <img src="https://raw.githubusercontent.com/aceberg/LightAlert/main/assets/logo.png" width="25" />
</a>LightAlert</h1>

Lightweight cron job monitoring 

- [Quick start](https://github.com/aceberg/lightalert#quick-start)
- [Config](https://github.com/aceberg/lightalert#config)
- [Options](https://github.com/aceberg/lightalert#options)
- [Thanks](https://github.com/aceberg/lightalert#thanks)


![Screenshot](https://raw.githubusercontent.com/aceberg/LightAlert/main/assets/Screenshot%202023-02-01%20at%2014-36-39%20LightAlert.png)
![Screenshot1](https://raw.githubusercontent.com/aceberg/LightAlert/main/assets/Screenshot%202023-02-01%20at%2014-34-35%20LightAlert.png)

## Quick start

```sh
docker run --name lightalert \
-e "TZ=Asia/Novosibirsk" \
-v ~/.dockerdata/LightAlert:/data/LightAlert \
-p 8846:8846 \
aceberg/lightalert
```
Or use [docker-compose.yml](docker-compose.yml)

## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| DB        | Path to Database | /data/LightAlert/sqlite.db |
| HOST | Listen address | 0.0.0.0 |
| PORT   | Port for web GUI | 8846 |
| THEME | Any theme name from https://bootswatch.com in lowcase | superhero |
| SHOW | How many lines to show on index page | 25 |
| YAMLPATH | Path to hosts yaml file | /data/LightAlert/hosts.yaml |
| TZ | Set your timezone for correct time | "" |

## Options

| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -c | Path to config file | /data/LightAlert/config.yaml | 
| -d | Path to SQLite DB file | /data/LightAlert/sqlite.db | 
| -h | Path to hosts yaml file | /data/LightAlert/hosts.yaml | 

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/LightAlert/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Favicon and logo: [Bell icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/bell)