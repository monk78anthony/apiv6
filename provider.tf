# ------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ------------------------------------------------------------------------------------------------------------------

provider "aws" {
  region     = var.region
  access_key = var.access_key
  secret_key = var.secret_key
}
