FROM centos:7.9.2009

# [must] maintainer label
LABEL "maintainer"="xiao_hua_ding@sina.cn"

# [must] timezone must be Asia/Shanghai (CST)
ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

RUN yum install -y telnet; yum clean all
RUN yum install -y tcpdump; yum clean all
RUN yum install -y wget; yum clean all
RUN yum install -y net-tools; yum clean all
RUN yum install -y bind-utils; yum clean all

RUN yum install -y sysstat; yum clean all
RUN yum install -y lsof; yum clean all
RUN yum install -y lrzsz; yum clean all
RUN yum install -y unzip; yum clean all

RUN sed -i 's/^ssl/#ssl/g' /etc/yum.conf

VOLUME /app/log
VOLUME /app/conf

# [suggest] working dir is app
ADD . /app
WORKDIR /app

# [suggest] declare exposed port
EXPOSE 10001
#ENV LD_LIBRARY_PATH lib:$LD_LIBRARY_PATH

# [suggest] use command array
# ENTRYPOINT指定默认的运行命令
ENTRYPOINT ["/bin/bash","-c"]
# CMD指定默认的运行参数
# CMD ["start.sh param1 param2"]
CMD ["start.sh"]
