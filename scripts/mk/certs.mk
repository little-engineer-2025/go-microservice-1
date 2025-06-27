##
# Helper rules to generate a developer self-signed
# certificate.
# Do not forget to add SSL_CERT_DIR when starting
# your microservice pointing out to your repository
# to serve requests using the temporary certificate.
#
# Note your system should resolve the value of SERVERNAME
# to 127.0.0.1 on development environments.
#
# see: man openssl-x509
##
CERT_DIR ?= $(PROJECT_DIR)/secrets
CERT_SERVERNAME ?= example.com
CERT_ALT_NAME ?= DNS:$(CERT_SERVERNAME),DNS:*.$(CERT_SERVERNAME),IP:127.0.0.1

.PHONY: generate-cert
generate-cert:  $(CERT_DIR)/$(CERT_SERVERNAME).key $(CERT_DIR)/$(CERT_SERVERNAME).crt  ## Generate a local certificate for development

$(CERT_DIR)/$(CERT_SERVERNAME).key $(CERT_DIR)/$(CERT_SERVERNAME).crt:
	openssl x509 \
		-new \
		-set_issuer "/CN=$(CERT_SERVERNAME)" \
		-subj "/CN=$(CERT_SERVERNAME)" \
		-not_before "today" \
		-days 30 \
		\
		-key rsa:4096 \
		-sha256 \
		-days 7 \
		-nodes \
		-keyout $(CERT_DIR)/$(CERT_SERVERNAME).key \
		-out $(CERT_DIR)/$(CERT_SERVERNAME).crt \
		-subj "/CN=$(CERT_SERVERNAME)" \
		-addext "subjectAltName=$(CERT_ALT_SERVERNAME)"

