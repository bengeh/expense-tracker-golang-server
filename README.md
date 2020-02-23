# expense-tracker-golang-server
How to start:
1. touch .env
2. write your details for mongodb for: MONGO_URL, DB_NAME, COllECTION_NAME
3. Run dockerscript to build a docker container
4. Use Postman/Curl etc. to do a call to localhost:41960
5. Should return a response.

6. run docker logs {container_id} --follow to tail the logs
7. docker images will show all images being built
8. docker ps shows all container
9. docker {container_id} stop -> stop a docker container
10. docker image rm {image_id}  --force -> remove a docker image