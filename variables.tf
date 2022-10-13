# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

variable "table_name" {
  description = "The name to use for all the cluster resources"
  type        = string
  default     = "apiv6-uat"
}

variable "read_capacity" {
  description = "The minimum number of EC2 Instances in the ASG"
  type        = number
  default     = 10
}

variable "write_capacity" {
  description = "The maximum number of EC2 Instances in the ASG"
  type        = number
  default     = 10
}
