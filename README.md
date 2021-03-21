1. Create a local mongo using docker
docker-compose up -d

2. Set env for mongo URL
export MONGO_URI_FOR_GO_CRUD=mongodb://localhost:27017

3. Verify that the above env is set correctly. 
echo $MONGO_URI_FOR_GO_CRUD