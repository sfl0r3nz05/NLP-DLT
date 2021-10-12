FROM node:11-alpine
WORKDIR /app
COPY . .
RUN npm install
CMD [ "npm", "start" ]

#FROM node:10.17.0-jessie
#RUN mkdir -p /app/client
#WORKDIR /app/client
#COPY package.json /app/client
#COPY package-lock.json /app/client
#RUN npm install
#COPY . /app/client
#CMD [ "npm", "start" ]

#FROM sflorenz05/contrack:frontend
#docker build -t frontend .
#docker tag 434a4a59dfa3 sflorenz05/contrack:frontend
#docker push sflorenz05/contrack:frontend
#docker ps
#docker commit d951e9d3eda9 sflorenz05/contrack:cadvisor
#docker push docker sflorenz05/contrack:cadvisor