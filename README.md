# Axobase

A web-server for the MDIBL Axobase resource.

## To run the service on a virtual machine:
    clone the repo
    cd <path_to>/axobase
    docker compose up (or docker-compose up)

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        HTML templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites