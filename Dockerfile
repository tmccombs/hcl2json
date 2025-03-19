FROM scratch
LABEL maintainer="Thayne McCombs <https://github.com/tmccombs>"
COPY hcl2json /hcl2json
ENTRYPOINT [ "/hcl2json" ]
