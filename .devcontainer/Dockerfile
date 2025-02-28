FROM fluxrm/flux-sched:jammy

LABEL maintainer="Vanessasaurus <@vsoch>"

# Match the default user id for a single system so we aren't root
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000
ENV USERNAME=${USERNAME}
ENV USER_UID=${USER_UID}
ENV USER_GID=${USER_GID}
USER root
RUN apt-get update && apt-get install -y jq less bats gnupg2 protobuf-compiler

# flux removed czmq
RUN /bin/bash -c "echo \"deb http://download.opensuse.org/repositories/network:/messaging:/zeromq:/release-stable/Debian_9.0/ ./\" >> /etc/apt/sources.list && \
    wget https://download.opensuse.org/repositories/network:/messaging:/zeromq:/release-stable/Debian_9.0/Release.key -O- | sudo apt-key add && \
    apt-get install -y libczmq-dev"

# Assuming installing to /usr/local
ENV LD_LIBRARY_PATH=/usr/local/lib

# Install Go (if Go bindings desired)
RUN wget https://go.dev/dl/go1.22.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.6.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin:/home/vscode/go/bin

# Add the group and user that match our ids
RUN groupadd -g ${USER_GID} ${USERNAME} && \
    adduser --disabled-password --uid ${USER_UID} --gid ${USER_GID} --gecos "" ${USERNAME} && \
    echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers
USER $USERNAME
