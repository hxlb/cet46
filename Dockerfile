FROM alpine
ADD cet46-srv /cet46-srv
ENTRYPOINT [ "/cet46-srv" ]
