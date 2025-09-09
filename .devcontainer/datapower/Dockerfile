FROM icr.io/cpopen/datapower/datapower-limited:10.6.0.6
ENV  DATAPOWER_ACCEPT_LICENSE=true \
     DATAPOWER_WORKER_THREADS=2 \
     DATAPOWER_INTERACTIVE=true

COPY config /opt/ibm/datapower/drouter/config
COPY local /opt/ibm/datapower/drouter/local
COPY cert /opt/ibm/datapower/root/secure/usrcerts
USER root
RUN chown -R drouter:root /opt/ibm/datapower/drouter/config \
                          /opt/ibm/datapower/drouter/local \
                          /opt/ibm/datapower/root/secure/usrcerts
RUN set-user drouter

USER drouter

EXPOSE 5554
EXPOSE 9090
