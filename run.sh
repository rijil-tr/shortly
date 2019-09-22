docker rm -f shortly
docker run -it  --name shortly -p 8080:8080 -e DEBUG="True" shortly