FROM allinurl/goaccess

ADD goaccess.conf /

ENV LANG="zh_CN.UTF-8"

ADD ./goaccess-cron /etc/cron.d/
RUN chmod 0644 /etc/cron.d/goaccess-cron
RUN mkdir -p /var/spool/cron/crontabs
RUN crontab /etc/cron.d/goaccess-cron
RUN mkdir -p /var/log/goaccess && touch /var/log/goaccess/cron.log

ENTRYPOINT ["crond", "-f"]
