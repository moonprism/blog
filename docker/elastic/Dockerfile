FROM elasticsearch:5.6.16

RUN wget https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v5.6.16/elasticsearch-analysis-ik-5.6.16.zip
ENV ikZip=elasticsearch-analysis-ik-5.6.16.zip
#COPY ${ikZip} ./
ENV pluginPath=/usr/share/elasticsearch/plugins

RUN cp ${ikZip} /tmp
RUN unzip /tmp/${ikZip} -d /tmp
RUN mv /tmp/elasticsearch ${pluginPath}/ik