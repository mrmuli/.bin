provider "aws" {
    access_key = "secretsauce"
    secret_key = "casper"
    region     = "${var.region}"
  
}
resource "aws_instance" "kilima_one" {
    ami           = "${var.ami_id}"
    instance_type = "${var.instance_type}"
  
}
resource "aws_instance" "kilima_two" {
    ami           = "${var.ami_id}"
    instance_type = "${var.instance_type}"
  
}