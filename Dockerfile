FROM alpine:3.2
ADD twitter-srv /twitter-srv
ENTRYPOINT [ "/twitter-srv" ]
