FROM nginx:1.19

ARG COMMIT_SHA=nogitsha

COPY index.html /usr/share/nginx/html
RUN sed -i s/%GIT_SHA%/${COMMIT_SHA}/ /usr/share/nginx/html/index.html
