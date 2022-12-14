# ---------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# --------------------------------------------------------------------------------------------------------

terraform {
  required_providers {
    aws = {
      source  = "aws"
      version = "~> 4.3"
    }
  }
      
  cloud {
    organization = "northernelephant"

    workspaces {
      name = "apiv6"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_dynamodb_table" "apiv6-uat" { 
   name = "apiv6-uat" 
   billing_mode = "PROVISIONED" 
   read_capacity = "10" 
   write_capacity = "10"
   hash_key = "band"
   range_key = "title"
   
   attribute { 
      name = "band" 
      type = "S" 
   } 

   attribute {
    name = "title"
    type = "S"
  }
  
   ttl { 
     enabled = true
     attribute_name = "expiryPeriod"  
   }

   point_in_time_recovery { enabled = false } 
   server_side_encryption { enabled = true } 
   
   lifecycle { ignore_changes = [ write_capacity, read_capacity ] }
}

output "id" {
  value       = aws_dynamodb_table.apiv6-uat.id
  description = "The domain name of the load balancer"
}

output "arn" {
  value       = aws_dynamodb_table.apiv6-uat.arn
  description = "The domain name of the load balancer"
}
