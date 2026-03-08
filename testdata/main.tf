# AWS VPC Resource
resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "xentic-core-vpc"
  }
}

# AWS Subnet public
resource "aws_subnet" "public" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = true

  tags = {
    Name = "public-subnet"
  }
}

# Dummy Module
module "database" {
  source = "./modules/db"

  vpc_id   = aws_vpc.main.id
  subnet   = aws_subnet.public.id
  
  # The following should be redacted
  password = "SuperSecretPassword123!"
  api_key  = "AKIAJ4Y73QABCDEXAMPLE"
}
