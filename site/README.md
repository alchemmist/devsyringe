# Devsyringe Landing

This is a simple static HTML landing page for Devsyringe. No build step.

## Local preview

```sh
cd site
python3 -m http.server 5143
```

Then open `http://localhost:5143`.

## Docker

```sh
cd site
make
```

This builds the image and serves the page on port `5143`.
