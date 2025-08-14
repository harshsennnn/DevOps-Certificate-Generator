# create a s3 bucket for terraform state

resource "aws_s3_bucket" "my_bucket" {
  bucket = "my-terraform-state-bucket_for_devops_url_2_qr"

  versioning {
    enabled = true
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }

}

#
# # create a dynamodb table for terraform state locking
# resource "aws_dynamodb_table" "state_lock" {
#   name = "terraform-state-lock"
#   billing_mode = "PAY_PER_REQUEST"
#   hash_key = "LockID"
#   attribute {
#     name = "LockID"
#     type = "S"
#   }
# }