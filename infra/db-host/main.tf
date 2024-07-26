provider "aws" {
  region  = "us-east-1"
  profile = "default"
}

# Use existing key pair for connecting (SSH) to instance
data "aws_key_pair" "personal_vm" {
  key_name = "cmu95737"
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-amd64-server-*"]
  }

  owners = ["099720109477"]
}

resource "aws_instance" "db_host" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "i4i.large"

  key_name = data.aws_key_pair.personal_vm.key_name
  user_data = file("${path.module}/user_data.sh")

  tags = {
    Name = "cmu95737-db-host"
  }
}
