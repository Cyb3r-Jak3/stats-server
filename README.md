# Stats Server

Golang server for collecting error data


### Endpoints

- `/report/404`  
  Input (JSON):
  - host: string
  - ip: string
  - method: string
  - time: int (epoch)

### Required Variables

- `DATABASE_URL`: POSTGRESQL database connection string