FROM nginx:alpine

COPY index.html /usr/share/nginx/html/index.html
COPY style.css /usr/share/nginx/html/style.css
COPY logo.svg /usr/share/nginx/html/logo.svg
COPY demo2.gif /usr/share/nginx/html/demo2.gif
COPY favicon.ico /usr/share/nginx/html/favicon.ico
COPY apple-touch-icon.png /usr/share/nginx/html/apple-touch-icon.png
COPY site.webmanifest /usr/share/nginx/html/site.webmanifest
COPY web-app-manifest-192x192.png /usr/share/nginx/html/web-app-manifest-192x192.png
COPY web-app-manifest-512x512.png /usr/share/nginx/html/web-app-manifest-512x512.png
COPY robots.txt /usr/share/nginx/html/robots.txt

EXPOSE 5143
CMD ["nginx", "-g", "daemon off;"]
