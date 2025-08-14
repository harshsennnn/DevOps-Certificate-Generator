terraform {
  backend "s3" {
    bucket         = "my-terraform-state-bucket_for_devops_url_2_qr"
    key            = "global/mystatefile/terraform.tfstate"
    region         = "ap-south-1"
    dynamodb_table = "terraform-state-lock"
    encrypt        = true

  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}


# Configure the AWS Provider
provider "aws" {
  region = "ap-south-1"
}