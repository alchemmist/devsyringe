FROM nginx:alpine

COPY index.html /usr/share/nginx/html/index.html
COPY style.css /usr/share/nginx/html/style.css
COPY logo.svg /usr/share/nginx/html/logo.svg
COPY demo2.gif /usr/share/nginx/html/demo2.gif
COPY favicon.ico /usr/share/nginx/html/favicon.ico
COPY robots.txt /usr/share/nginx/html/robots.txt

EXPOSE 5143
CMD ["nginx", "-g", "daemon off;"]
