FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD server /
CMD ["/server"]
