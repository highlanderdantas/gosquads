#env GOOS=linux GOARCH=amd64 go build
#docker build -t mercurius:gosquads .
#docker run -p 8080:8080 -d mercurius:gosquads

FROM scratch

ADD gosquads /
ADD conf/ /conf
ADD public/ /public
ADD locale/ /locale

CMD [ "/gosquads" ]