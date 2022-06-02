FROM projects.registry.vmware.com/tanzu_isv_engineering/mkpcli
COPY bin/check /opt/resource/check
COPY bin/in /opt/resource/in
COPY bin/out /opt/resource/out
WORKDIR /opt/resource
