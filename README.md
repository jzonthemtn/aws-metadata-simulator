# AWS Instance Metadata Simulator

This project provides an incomplete capability of simulating AWS EC2 [instance metadata](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html). It was created to allow for testing of applications that utilize EC2 instance metadata in non-AWS environments.

`go get -u github.com/mtnfog/aws-metadata-emulator`

To run:

If specific values for the instance metadata are desired set those values in `metadata.toml` then run it:

`go run main.go`

To use a different file as the configuration give the filename as a command line argument:

`go run main.go other.toml`

To redirect traffic:

`iptables -t nat -A OUTPUT -p tcp -d 169.254.169.254 --dport 80 -j DNAT --to-destination 127.0.0.1:8080`

To use:

`curl -X GET http://127.0.0.1:8080/latest/meta-data/ami-id`

## Summary of Usage

```
yum install git golang
mkdir ~/go
export GOPATH=~/go
go get -u github.com/mtnfog/aws-metadata-emulator
cd ~/go/bin
wget https://raw.githubusercontent.com/mtnfog/aws-metadata-simulator/master/metadata.toml
# Edit the metadata.toml to have the values you need.
sudo iptables -t nat -A OUTPUT -p tcp -d 169.254.169.254 --dport 80 -j DNAT --to-destination 127.0.0.1:8080
./aws-meta-simulator
```
