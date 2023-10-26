FROM golang:1.21.1

RUN curl -L https://foundry.paradigm.xyz | bash && . /root/.bashrc && foundryup
