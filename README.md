# Axobase

A web-server for the MDIBL Axobase resource.

# Making changes

*Note: This app is built on x86 architecture, which means that building on ARM architecture will fail. \
To get around this, make and test changes on Random or another x86 machine. \
Axobase is hosted on AWS App Runner, which deploys the AWS ECR image. \

Before building the new local image, remove old axobase images using 'docker images | grep "axo"' \
Run 'docker rmi -f  <image name>' to remove the old images. \

Pull or clone the GitHub repo and make desired changes. \
Then use 'docker compose up' to build the new image. \
You can view the local changes via ssh tunnel to port 9100 of the server that you're testing on. \
Example: ssh -L 9100:localhost:9100 username@random.mdibl.org \

Next, tag and push the image to our AWS ECR repo: \
First, retrieve AWS credentials with \
'aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 012870262837.dkr.ecr.us-east-1.amazonaws.com'

Tag the newly built image using 'docker tag axobase-axobase-app:latest 012870262837.dkr.ecr.us-east-1.amazonaws.com/axobase:latest'

Push the newly tagged image to ECR using 'docker push 012870262837.dkr.ecr.us-east-1.amazonaws.com/axobase:latest'

The App Runner service automatically detects the updated image and redeploys the service.
This can be verified by navigating to the AWS AppRunner service in the console and/or waiting a few minutes and reloading axobase.org

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
