FROM node:10.9
WORKDIR /home/node/app/
COPY run.js .
ENTRYPOINT ["node","run.js"]
EXPOSE 7070
