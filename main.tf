terraform {
  required_version = "1.3.2"
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
   name = var.table_name 
   billing_mode = "PROVISIONED" 
   read_capacity = var.read_capacity 
   write_capacity = var.write_capacity
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
