FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD animalgame /
CMD ["/animalgame"]
