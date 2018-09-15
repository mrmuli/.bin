
module "aws" {
  source = "./aws"
  ami = "${var.ami}" # experimental
  
}
resource "docker_image" "kilimanjaro_image" {
    name = "${var.image_name}"
}
resource "docker_container" "kilimanjaro" {
    name  = "${var.container_name}" 
    image = "${var.image_name}"
    ports {
        internal = "${var.internal_port}"
        external = "${var.external_port}"
    }
}