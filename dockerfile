FROM nimmis/alpine-micro
RUN apk update && apk upgrade

ENV TERRAFORM_VERSION=0.13.7

# Install kubernetes CLI tool
RUN apk update && \
    apk add --update bash curl python3 py3-pip && \
    curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && \
    mv kubectl /usr/local/bin/kubectl && \
    chmod +x /usr/local/bin/kubectl && \
    rm -rf /var/cache/apk/*

# Install terraform
RUN curl -LO https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    rm -f terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    chmod +x terraform && \
    mv terraform /usr/local/bin

WORKDIR /

ADD config $HOME/.kube/config

CMD [ "tail", "-f", "/dev/null" ]