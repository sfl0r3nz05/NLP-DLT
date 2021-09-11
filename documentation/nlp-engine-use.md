# How to use the NLP-Engine 😎

## Set up 🙂
1. Clone the repository
2. Build a docker image:
    - `cd ~/nlp-dlt/nlp`
    - `docker build -t nlp-engine .`
3. To verify the docker image should be used the command `docker images`
    <img src="https://github.com/sfl0r3nz05/nlp-dlt/blob/sentencelvl/documentation/images/dockerVerification.png">
3. Create an environmental variables file based on .env.example (~/nlp-dlt/network)
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