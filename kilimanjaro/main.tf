provider "aws" {
    access_key = "secretsauce"
    secret_key = "saucetisecret"
    region     = "us-east-2"
  
}

resource "aws_instance" "kilima_one" {
    ami           = "ami-2757f631"
    instance_type = "t2.micro"
  
}
resource "aws_instance" "kilima_two" {
    ami           = "ami-2757f631"
    instance_type = "t2.micro"
  
}
resource "docker_image" "kilimanjaro_image" {
    name = "josephmuli/simple_nginx:0.0.2"
}

resource "docker_container" "kilimanjaro" {
    name  = "kilimanjaro_container" 
    image = "${docker_image.kilimanjaro_image.latest}"
    ports {
        internal = "8080"
        external = "8080"
    }
}
