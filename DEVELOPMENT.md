# How to use the development environment

Pre-requisites:

* Follow Getting Started steps at `README.md` file.

## Accounts

In order to use the ephemeral development environment you need several accounts.

* Github
* Quay
* RedHat Registry

See also:
<https://consoledot.pages.redhat.com/docs/dev/getting-started/ephemeral/getting-started-with-ees.html#_join_redhatinsights>

### Github

To get access to the ephemeral development environment you need to be a member
of the RedHatInsights organization at Github.

To get your account added to the RedHatInsights organization use the "add user
to github" form at:

<https://source.redhat.com/groups/public/consoledot/consoledot_wiki/console_dot_requests_submission_process#submission-process-for-operational-requests>

### Quay.io

Go to <https://quay.io/> and login with your Red Hat user account.

* Create a new repository.
* Go to <https://quay.io/repository/<rh> user>/<repository>?tab=settings
* Create a new robot account and give the account write permissions to your brand new image repository.
* Create the secrets file: `cp ./scripts/mk/private.example.mk secrets/private.mk`
* Fill out the login and token in `secrets/private.mk`

### Install python dependencies

    make .venv
    source .venv/bin/activate
    pip install -r requirements-dev.txt

## Deploying to ephemeral

    make ephemeral-login
    make ephemeral-namespace-create
    make ephemeral-deploy

    # If correct image was already built and pushed,
    # set EPHEMERAL_NO_BUILD=y to skip these steps.
    make ephemeral-deploy EPHEMERAL_NO_BUILD=y

## Launching request against the API

From VSCode or any IDE which support the .http files, we could open
the scripts at `scripts/http/public.http` and launch the http commands
found there.

> You will need yq tool installed to read values in Makefile from `configs/config.yaml` file.
> See: <https://github.com/mikefarah/yq#install>

### Locally

To test our API locally, we can start the service by `make compose-up run` and launching
the test script below:

    ```sh
    ./test/scripts/local-todos-create.sh
    ```

You can use the `.http` files.
