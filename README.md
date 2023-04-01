# Go-Sample-Project

Try Go with Fiber and Wire DI.

Service check all url from service.json (also expose self health-check) and send alert notification to telegram chat if down

## Usage

```bash
   docker-compose up --build checker-dev // for develop
   
   docker-compose up --build checker // for prod
```

## Docs:

- [Fiber](https://gofiber.io/)
- [Wire](https://github.com/google/wire)
- [Testify](https://github.com/stretchr/testify)

## TODO: 
https://projectbook.code.brettchalupa.com/web-apps/uptime-checker.html
