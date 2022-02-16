# Axobase

A web-server for the MDIBL Axobase resource.

## To run the service on a virtual machine:
    clone the repo
    cd <path_to>/axobase
    docker compose up (or docker-compose up)

The above command will expose the service on localhost e.g., http://127.0.0.1:9100. This does not 
provide external access to the application from the web. To gain access to the server from the web,
the machine must be configured with proxy such as Apache or Nginx which can forward all traffic to
127.0.01:9100. Alternatively, the Revel server itself could be run as root and configured to listen on 
port 80. This is not advised. 

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