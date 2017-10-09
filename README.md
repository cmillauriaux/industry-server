# Industry Golang server
TODO : Description

## Configuration
Industry Golang Server require an Amazon WebServices DynamoDb database. You must provide region and credentiels to run it.
```
export AWS_REGION=us-west-2
export AWS_ACCESS_KEY_ID=YOUR_AKID
export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY
export AWS_SESSION_TOKEN=TOKEN
```

## Deploy to AWS Lambda (optional)
You can deploy the server on AWS Lambda :
### Usage
Setup and deploy a new project called `your-app`:

```bash
cd you-app
make DOTENV=.env.example dotenv

```
* fill in and correct any of the variables in .env
* replace `WORKDIR` in .env with `/go/src/your-path/your-app`

```bash
make test build deploy
```

should get the same payload back with status `200`

```bash
make remove
```
