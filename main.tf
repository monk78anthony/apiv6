variable "region" {}
variable "access_key" {}
variable "secret_key" {}
variable "table_name" {}
variable "read_capacity" {}
variable "write_capacity" {}

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

output "id" {
  value       = aws_dynamodb_table.apiv6-uat.id
  description = "The domain name of the load balancer"
}

output "arn" {
  value       = aws_dynamodb_table.apiv6-uat.arn
  description = "The domain name of the load balancer"
}
