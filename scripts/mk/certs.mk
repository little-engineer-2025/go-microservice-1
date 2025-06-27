##
# Helper rules to generate a developer self-signed
# certificate.
# Do not forget to add SSL_CERT_DIR when starting
# your microservice pointing out to your repository
# to serve requests using the temporary certificate.
#
# Note your system should resolve the value of SERVERNAME
# to 127.0.0.1 on development environments.
##
SSL_CERT_DIR ?= $(PROJECT_DIR)
SERVERNAME ?= example.com

.PHONY: gen-local-cert
gen-local-cert:  $(SERVERNAME).key $(SERVERNAME).crt  ## Generate a local certificate for development

$(SERVERNAME).key $(SERVERNAME).crt:
	openssl req -x509 -newkey rsa:4096 -sha256 -days 7 -nodes -keyout local.key -out local.crt -subj "/CN=$(SERVERNAME) -addext "subjectAltName=DNS:$(SERVERNAME),DNS:*.$(SERVERNAME),IP:127.0.0.1"
