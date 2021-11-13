FROM node:13-alpine
WORKDIR /app/src
COPY src/ .
RUN npm install
CMD [ "npm", "start" ]
#FROM node:10.16.0-alpine
#RUN mkdir -p /opt/app
#WORKDIR /opt/app
#RUN adduser --disabled-password app
#COPY src/ .
#RUN chown -R app:app /opt/app
#USER app
#RUN npm install
#CMD [ "npm", "run", "start" ]
#FROM sflorenz05/contrack:gateway
#docker build -t backend .
#docker tag 16fe245f9f4a sflorenz05/contrack:gateway
#docker push sflorenz05/contrack:gateway