# AWS Instance Metadata Simulator

This project provides an incomplete capability of simulating AWS EC2 instance metadata. It was created to allow for testing of applications that utilize EC2 instance metadata in non-AWS environments.

`go get -u github.com/mtnfog/aws-metadata-emulator`

To run:

`go run main.go`

To redirect traffic:

`iptables -t nat -A PREROUTING -d 169.254.169.254 -p tcp --dport 80 --syn -j REDIRECT --to-port 8080`

To use:

`curl -X GET http://127.0.0.1:8080/latest/meta-data/ami-id`