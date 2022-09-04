# FROM quay.io/keycloak/keycloak:legacy
FROM sleighzy/keycloak
ARG ADMIN_USERNAME
ARG ADMIN_PASSWORD
RUN /opt/jboss/keycloak/bin/add-user-keycloak.sh -u $ADMIN_USERNAME -p $ADMIN_PASSWORD