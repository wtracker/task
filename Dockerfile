FROM alpine
ADD task /task
ENTRYPOINT [ "/task" ]
