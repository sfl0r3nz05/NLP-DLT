# How to use the NLP-Engine 😎

1. Go to the project directory:
    - `cd $GOPATH/src/github.com/nlp-dlt/nlp-engine`

2. Build a docker image:
    - `docker build -t nlp-engine .`

2. Verify the `docker image` built through the command `docker images`

    <img src="https://github.com/sfl0r3nz05/nlp-dlt/blob/sentencelvl/documentation/images/dockerVerification.png">

3. Create the environmental variables file based on .env.example ($GOPATH/src/github.com/nlp-dlt/network/nlp-engine)

4. Obtain access keys from AWS E.g.:
    <img src="https://github.com/sfl0r3nz05/nlp-dlt/blob/sentencelvl/documentation/images/accessKey.png">
    - Update the environment variable `AWS_ACCESS_KEY_ID`
    - Update the environment variable `AWS_SECRET_ACCESS_KEY`
    - Update the environment variable `AWS_SESSION_TOKEN`
    - Update the path of PDF files which contains Roaming Agreements
        - Default place: `~/nlp-dlt/nlp/input`
    - Update the path of JSON files
        - Default place: `~/nlp-dlt/nlp/output`

## Deploy the NLP-Engine 🙂
1. `cd ~/nlp-dlt/network`
2. Start: `docker-compose up -d`
3. Stop: `docker-compose stop`
4. Down: `docker-compose down`