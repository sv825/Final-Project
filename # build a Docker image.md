# build a Docker image
build:
    docker build -t voterapiredis .
    docker build -t voteapi . 
    docker build -t pollapi .