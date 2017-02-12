# AWS Instance Metadata Simulator

This project provides an incomplete capability of simulating AWS EC2 [instance metadata](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html). It was created to allow for testing of applications that utilize EC2 instance metadata in non-AWS environments.

`go get -u github.com/mtnfog/aws-metadata-emulator`

To run:

If specific values for the instance metadata are desired set those values in `metadata.toml` then run it:

`go run main.go`

To redirect traffic:

`iptables -t nat -A PREROUTING -d 169.254.169.254 -p tcp --dport 80 --syn -j REDIRECT --to-port 8080`

To use:

`curl -X GET http://127.0.0.1:8080/latest/meta-data/ami-id`
