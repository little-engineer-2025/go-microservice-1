##
# Helper rules to generate a developer self-signed certificate.
#
# Do not forget to add SSL_CERT_DIR when starting your microservice pointing out
# to your repository to serve requests using the temporary certificate.
#
# Note your system should resolve the value of CERT_SERVERNAME to 127.0.0.1 on
# development environments.
#
# see:
#   man openssl-x509
#   man x509v3_config
##
CERT_DIR ?= $(PROJECT_DIR)/secrets/certs
CERT_CNF_TEMPLATE ?= configs/certificate.cnf
CERT_DN ?= example.com
# CERT_SERVERNAME ?= something.$(CERT_DN)
CERT_IP ?= 127.0.0.1
CERT_ALT_NAME ?= DNS:$(CERT_SERVERNAME),IP:$(CERT_IP)

CERT_SERVERNAME_LIST ?= api.$(CERT_DN) db.$(CERT_DN) minio.$(CERT_DN) redis.$(CERT_DN)

.PHONY: cert-gen-all
cert-gen-all:  ## Generate all the local certificates for the service and infrastructure services
	@[ "$(CERT_SERVERNAME_LIST)" != "" ] || { echo "error: try 'make cert-gen CERT_SERVERNAME_LIST=\"api.example.com\"': CERT_SERVERNAME_LIST is empty"; exit 1; }
	for item in $(CERT_SERVERNAME_LIST); do \
		$(MAKE) cert-gen CERT_SERVERNAME="$$item" || { echo "error: cert-gen-all for $$item"; exit 1; }; \
	done

.PHONY: cert-gen
cert-gen: $(CERT_DIR) $(CERT_DIR)/$(CERT_SERVERNAME).key $(CERT_DIR)/$(CERT_SERVERNAME).crt  ## Generate a local certificate for development for CERT_SERVERNAME fqdn

$(CERT_DIR):
	mkdir -p "$(CERT_DIR)"

# $(CERT_DIR)/%.key $(CERT_DIR)/%.crt: $(CERT_DIR)/%.cnf
#	exit 1
#	openssl req -x509 -nodes -days 7 -newkey rsa:2048 -keyout "$(CERT_DIR)/$(CERT_SERVERNAME).key" -out "$(CERT_DIR)/$(CERT_SERVERNAME).crt" -config "$(CERT_DIR)/$(CERT_SERVERNAME).cnf"

$(CERT_DIR)/%.key $(CERT_DIR)/%.csr $(CERT_DIR)/%.key: $(CERT_DIR)/%.cnf
	openssl req -new \
		-config "$(CERT_DIR)/$(CERT_SERVERNAME).cnf" \
		-keyout "$(CERT_DIR)/$(CERT_SERVERNAME).key" \
		-out "$(CERT_DIR)/$(CERT_SERVERNAME).csr" \
		-nodes

$(CERT_DIR)/%.crt: $(CERT_DIR)/%.key $(CERT_DIR)/%.csr
	openssl x509 -req \
		-config "$(CERT_DIR)/$(CERT_SERVERNAME).cnf" \
		-days 365 \
		-in "$(CERT_DIR)/$(CERT_SERVERNAME).csr" \
		-signkey "$(CERT_DIR)/$(CERT_SERVERNAME).key" \
		-out "$(CERT_DIR)/$(CERT_SERVERNAME).crt"

$(CERT_DIR)/%.cnf:
	cat "$(CERT_CNF_TEMPLATE)" \
		| sed -e "r/commonName =/commonName = $(CERT_SERVERNAME)/g" \
		| sed -e "r/countryName =/countryName = US/g" \
		| sed -e "r/stateOrProvinceName =/stateOrProvinceName = California/g" \
		| sed -e "r/localityName =/localityName = San Francisco/g" \
		| sed -e "r/organizationName =/organizationName = My Company/g" \
		| sed -e "r/DNS.1 =/DNS.1 = $(CERT_SERVERNAME)/g" \
		| sed -e "r/IP.1 =/IP.1 = $(CERT_IP)/g" > "$(CERT_DIR)/$(CERT_SERVERNAME).cnf" \
		| tee "$@"

